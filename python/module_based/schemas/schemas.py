from pydantic import BaseModel


class StoredItem(BaseModel):
    id: str
    name: str


class PostRequest(BaseModel):
    name: str
