import asyncio
from typing import Any, Iterable, List

import polars as pl
from minio import Minio
from pydantic import BaseModel
from src.embeddings import Embeddings
from structlog import get_logger
from tiktoken.core import Encoding

logger = get_logger("utils")


class QueueMessage(BaseModel):
    lore_id: str
    bucket: str
    key: str
    n_rows: int


async def collect(items: List[List[Any]]) -> List[Any]:
    return [x for y in items for x in y]


def get_data_generator(
    client: Minio, bucket: str, key: str, output_path: str, n_rows: int
) -> Iterable[str]:
    client.fget_object(bucket, key, output_path)
    lf = pl.scan_parquet(output_path)
    for x in range(n_rows):
        df = lf.slice(x).collect()
        yield df["content"].to_list()[0]


async def batch_process(
    inputs: List[str], embeddings: Embeddings, batch_size: int
) -> List[List[float]]:
    tasks = []
    for i in range(0, len(inputs), batch_size):
        batch = inputs[i : i + batch_size]
        tasks.append(embeddings.embed(batch))

    results = await asyncio.gather(*tasks)
    flat_results = await collect(results)
    return flat_results


async def split_data(text: str, enc: Encoding, max_tokens: int) -> List[str]:
    tokens = enc.encode(text)
    content = []
    token_num = 0
    record = []
    for token in tokens:
        token_num += 1
        record.append(token)
        if token_num == max_tokens:
            content.append(enc.decode(record))
            token_num = 0
            record = []
    content.append(enc.decode(record))
    return content
