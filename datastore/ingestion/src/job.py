from git import Repo
from structlog import get_logger
import os
from typing import Iterable, List, Any
from embeddings import Embeddings
import asyncio
import tiktoken
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
):
    text = await split_data(content, encoder, 512)
    embeddings = await batch_process(text, embeddings, batch_size)


async def main():
    embeddings = Embeddings("http://127.0.0.1:3000")
    encoder = tiktoken.get_encoding("o200k_base")
    get_data("data")
    data_generator = get_data_generator("data")
    tasks = []
    for content in data_generator:
        tasks.append(process(content, encoder, 512, embeddings, 10))
    await asyncio.gather(*tasks)


asyncio.run(main())
