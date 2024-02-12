from fastapi import FastAPI

import uvicorn

from router.base import router

app = FastAPI()

app.include_router(router)

if __name__ == '__main__':
    uvicorn.run(app=app, port=8080)
