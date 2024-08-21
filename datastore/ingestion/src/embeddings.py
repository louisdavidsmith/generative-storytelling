import logging

import httpx
from structlog import get_logger
from tenacity import (before_sleep_log, retry, stop_after_attempt,
                      wait_exponential)

logger = get_logger("embeddings")


class Embeddings:
    def __init__(self, base_url):
        self.base_url = base_url
        self.client = httpx.AsyncClient(timeout=240)

    @retry(
        wait=wait_exponential(multiplier=1, min=4, max=10),
        stop=stop_after_attempt(4),
        before_sleep=before_sleep_log(logger, logging.INFO),
    )
    async def embed(self, inputs):
        data = {"inputs": inputs}
        url = f"{self.base_url}/embed"
        response = await self.client.post(url, json=data)
        response.raise_for_status()
        return response.json()

    async def close(self):
        await self.client.aclose()
