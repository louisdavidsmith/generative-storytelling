import argparse
import asyncio
import json
import os
from typing import Any, Coroutine, List, Tuple

import aio_pika
import tiktoken
from minio import Minio
from pydantic import ValidationError
from src.embeddings import Embeddings
from src.postgres_utils import get_conn, write_embeddings
from src.utils import QueueMessage, batch_process, get_data_generator, split_data
from structlog import get_logger
from tiktoken.core import Encoding

logger = get_logger("lore-ingestion")


async def process(
    content: str,
    encoder: Encoding,
    max_token_span: int,
    embeddings: Embeddings,
    batch_size: int,
    lore_id: str,
) -> Tuple[List[str], List[List[float]]]:
    text = await split_data(content, encoder, max_token_span)
    embedded_content = await batch_process(text, embeddings, batch_size)
    return text, embedded_content


async def main(args: argparse.Namespace):
    queue_name = args.queue_name

    connection = await aio_pika.connect_robust(args.amqp_url)
    channel = await connection.channel()
    queue = await channel.declare_queue(queue_name, durable=True)

    async def process_message(
        message: aio_pika.IncomingMessage,
    ):
        embeddings = Embeddings(args.embeddings_url)
        encoder = tiktoken.get_encoding("o200k_base")
        pool = await get_conn(
            args.host,
            args.dbname,
            args.user,
            args.password,
        )
        logger.info("PostgresPoolCreated")

        client = Minio(
            "minio:9000", access_key="minioadmin", secret_key="minioadmin", secure=False
        )
        async with message.process():
            payload = message.body.decode()
            payload = json.loads(payload)
            try:
                queue_message = QueueMessage(**payload)
            except ValidationError as exc:
                logger.info("ErrorReadingMessage", exc=exc)
                raise exc
            output_path = f"{queue_message.lore_id}.parquet"
            logger.info("MessageRecieved", data=queue_message)
            data_generator = get_data_generator(
                client,
                queue_message.bucket,
                queue_message.key,
                output_path,
                queue_message.n_rows,
            )
            tasks: List[Coroutine[Any, Any, Any]] = []
            for content in data_generator:
                tasks.append(
                    process(
                        content,
                        encoder,
                        args.max_tokens,
                        embeddings,
                        args.batch_size,
                        args.lore_id,
                    )
                )
            out = await asyncio.gather(*tasks)
            logger.info("DataScored")
            embeded_lore = []
            lore_content = []
            for batch in out:
                lore_content.extend(batch[0])
                embeded_lore.extend(batch[1])
            await write_embeddings(
                pool, embeded_lore, lore_content, queue_message.lore_id
            )
            logger.info("DataIngested")

    await queue.consume(process_message, auto_ack=True)

    logger.info("Waiting for messages...")
    await asyncio.Future()  # Run forever


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
    parser.add_argument("--amqp_url", type=str, default=os.getenv("AMQP_URL", ""))
    parser.add_argument(
        "--queue_name", type=str, default=os.getenv("QUEUE_NAME", "data_queue")
    )

    args = parser.parse_args()

    asyncio.run(main(args))
