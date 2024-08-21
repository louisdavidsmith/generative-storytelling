import httpx
from tenacity import retry, wait_exponential


class Embeddings:
    def __init__(self, base_url):
        self.base_url = base_url
        self.client = httpx.AsyncClient(timeout=240)

    @retry(wait=wait_exponential(multiplier=1, min=4, max=10))
    async def embed(self, inputs):
        data = {"inputs": inputs}
        url = f"{self.base_url}/embed"
        response = await self.client.post(url, json=data)
        response.raise_for_status()
        return response.json()

    async def close(self):
        await self.client.aclose()
