import uuid
from typing import List

from fastapi import APIRouter, HTTPException, status
from schemas.schemas import PostRequest, StoredItem

router = APIRouter()

data: List[StoredItem] = []


@router.get("/ping")
def handle_ping() -> str:
    return "pong"


@router.get("/")
def handle_get_all() -> List[StoredItem]:
    return data


@router.get("/{item_id}")
def handle_get_by_id(item_id: str) -> StoredItem:
    for item in data:
        if item.id == item_id:
            return item
    raise HTTPException(detail="not found", status_code=status.HTTP_404_NOT_FOUND)


@router.post("/")
def handle_item_create(item: PostRequest) -> StoredItem:
    resp = StoredItem(
        id=str(uuid.uuid4()),
        name=item.name,
    )
    data.append(resp)
    return resp
