from fastapi import FastAPI
from backend.middleware.logging_middleware import logging_middleware
from backend.api.books import router as books_router

app = FastAPI(
    title="Fastapi Digital Library API",
    description="A simple REST API with CRUD operations, error handling and data validation to manage books in a digital library",
    version="1.0.0",
)

app.middleware("http")(logging_middleware)
app.include_router(books_router)  

@app.get("/")
def root():
    return {"message": "Digital Library API is running"}