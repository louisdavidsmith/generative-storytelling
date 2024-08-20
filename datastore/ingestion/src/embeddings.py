import httpx


class Embeddings:
    def __init__(self, base_url):
        self.base_url = base_url
        self.client = httpx.AsyncClient(timeout=60)

    async def embed(self, inputs):
        data = {"inputs": inputs}
        url = f"{self.base_url}/embed"
        response = await self.client.post(url, json=data)
        response.raise_for_status()
        return response.json()

    async def tokenize(self, inputs):
        data = {"inputs": inputs}
        url = f"{self.base_url}/tokenize"
        response = await self.client.post(url, json=data)
        response.raise_for_status()
        return response.json()

    async def close(self):
        await self.client.aclose()
