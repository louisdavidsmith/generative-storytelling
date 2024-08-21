import argparse
import asyncio
import os
from typing import Any, Coroutine, Iterable, List

import tiktoken
from git import Repo
from src.embeddings import Embeddings
from src.postgres_utils import get_conn, write_embeddings
from structlog import get_logger
from tiktoken.core import Encoding

logger = get_logger("lore-ingestion")


async def collect(items: List[List[Any]]) -> List[Any]:
    return [x for y in items for x in y]


def get_data(path: str):
    if os.path.exists(path):
        logger.info("DataAlreadyPresent")
    else:
        Repo.clone_from("https://github.com/vilmibm/lovecraftcorpus.git", path)
        logger.info("DataRetrieved")


def get_data_generator(path: str, extension=".txt") -> Iterable[str]:
    files = os.listdir(path)
    text_files = [x for x in files if extension in x]
    for file_name in text_files:
        with open(f"{path}/{file_name}", "r", encoding="utf-8") as file:
            content = file.read()
            yield content


async def batch_process(inputs: List[str], embeddings: Embeddings, batch_size: int):
    tasks = []
    for i in range(0, len(inputs), batch_size):
        batch = inputs[i : i + batch_size]
        tasks.append(embeddings.embed(batch))

    results = await asyncio.gather(*tasks)
    flat_results = await collect(results)
    return flat_results


async def split_data(text: str, enc: Encoding, max_tokens: int) -> List[Any]:
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


async def process(
    content: str,
    encoder: Encoding,
    max_token_span: int,
    embeddings: Embeddings,
    batch_size: int,
    pool,
    lore_id: str,
):
    text = await split_data(content, encoder, 512)
    embeddings = await batch_process(text, embeddings, batch_size)
    await write_embeddings(pool, embeddings, text, lore_id)
    logger.info("DataIngested", n_records=len(text))


async def main(args: argparse.Namespace):
    embeddings = Embeddings(args.embeddings_url)
    encoder = tiktoken.get_encoding("o200k_base")
    pool = await get_conn(
        args.host,
        args.dbname,
        args.user,
        args.password,
    )
    get_data("data")
    data_generator = get_data_generator("data")
    tasks: List[Coroutine[Any, Any, Any]] = []
    for content in data_generator:
        tasks.append(
            process(
                content,
                encoder,
                args.max_tokens,
                embeddings,
                args.batch_size,
                pool,
                args.lore_id,
            )
        )
    await asyncio.gather(*tasks)
    logger.info("DataIngested")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Process data with embeddings.")
    parser.add_argument("--lore_id", type=str, default=os.getenv("LORE_ID", ""))
    parser.add_argument(
        "--embeddings_url", type=str, default=os.getenv("EMBEDDINGS_URL", "")
    )
    parser.add_argument("--host", type=str, default=os.getenv("HOST", ""))
    parser.add_argument("--dbname", type=str, default=os.getenv("DBNAME", ""))
    parser.add_argument("--user", type=str, default=os.getenv("USER", ""))
    parser.add_argument("--password", type=str, default=os.getenv("PASSWORD", ""))
    parser.add_argument(
        "--batch_size", type=int, default=int(os.getenv("BATCH_SIZE", "10"))
    )
    parser.add_argument(
        "--max_tokens", type=int, default=int(os.getenv("MAX_TOKENS", "512"))
    )

    args = parser.parse_args()

    asyncio.run(main(args))
