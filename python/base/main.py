import uuid
from typing import List

from fastapi import FastAPI, HTTPException, status

import uvicorn
from pydantic import BaseModel

app = FastAPI()


class StoredItem(BaseModel):
    id: str
    name: str


class PostRequest(BaseModel):
    name: str


data: List[StoredItem] = []


@app.get("/ping")
def handle_ping() -> str:
    return "pong"


@app.get("/")
def handle_get_all() -> List[StoredItem]:
    return data


@app.get("/{item_id}")
def handle_get_by_id(item_id: str) -> StoredItem:
    for item in data:
        if item.id == item_id:
            return item
    raise HTTPException(detail="not found", status_code=status.HTTP_404_NOT_FOUND)


@app.post("/")
def handle_item_create(item: PostRequest) -> StoredItem:
    resp = StoredItem(
        id=str(uuid.uuid4()),
        name=item.name,
    )
    data.append(resp)
    return resp


if __name__ == '__main__':
    uvicorn.run(app=app, port=8080)
