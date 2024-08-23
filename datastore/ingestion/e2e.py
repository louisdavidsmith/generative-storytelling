import asyncio

import aio_pika

import json


async def main() -> None:
    connection = await aio_pika.connect_robust(
        "amqp://guest:guest@127.0.0.1/",
    )

    async with connection:
        routing_key = "data_queue"

        message = json.dumps(
            {
                "bucket": "ingestion",
                "key": "staging/data.parquet",
                "n_rows": 67,
                "lore_id": "be846c26-8889-4ca6-9eab-cc867501ad22",
            }
        )
        message = aio_pika.Message(body=message.encode())
        channel = await connection.channel()
        await channel.default_exchange.publish(
            message,
            routing_key=routing_key,
        )


if __name__ == "__main__":
    asyncio.run(main())
