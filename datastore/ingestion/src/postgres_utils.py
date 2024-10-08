import asyncio
import json
import uuid
from typing import List

import psycopg_pool
from structlog import get_logger
from tenacity import retry, wait_exponential

logger = get_logger("ps-utils")


async def get_conn(
    host: str, dbname: str, user: str, password: str, max_size: int = 10
) -> psycopg_pool.AsyncConnectionPool:
    conn_info = f"host={host} dbname={dbname} user={user} password={password}"
    logger.info("ConnInfo", conn_info=conn_info)
    pool = psycopg_pool.AsyncConnectionPool(
        conninfo=conn_info, open=False, max_size=max_size
    )
    return pool


@retry(wait=wait_exponential(multiplier=1, min=4, max=10))
async def write_embeddings(
    pool: psycopg_pool.AsyncConnectionPool,
    embeddings: List[List[float]],
    content: List[str],
    lore_id: str,
):
    await pool.open()
    await pool.wait()
    query = """
        INSERT INTO lore(lore_id, embedding, content)
        VALUES(%s, %s, %s)
        ON CONFLICT (lore_id, embedding) DO NOTHING;
    """
    data = [
        (lore_id, json.dumps(embedding), record)
        for embedding, record in zip(embeddings, content)
    ]
    async with pool.connection() as conn:
        async with conn.cursor() as cursor:
            await cursor.executemany(query, data, returning=False)


async def ping(pool: psycopg_pool.AsyncConnectionPool):
    await pool.open()
    await pool.wait()
    async with pool.connection() as conn:
        async with conn.cursor() as cursor:
            await cursor.execute("SELECT 1")
            results = await cursor.fetchall()
            print(results)
    await pool.close()


if __name__ == "__main__":

    async def main():
        pool = await get_conn("localhost", "lovecraft", "lovecraft")
        await write_embeddings(pool, [[0.3, 5, 0.3]], ["Hello"], str(uuid.uuid4()))

    asyncio.run(main())
