from fastapi import APIRouter, HTTPException
from backend.model.book_model import Book
from typing import List
router = APIRouter(
    prefix="/books",
    tags=["Library"],
)

books_db: List[Book] = []

@router.get("/",
    summary="Get all books in the library",
    description="Retrieves the complete list of books available in the digital library",
    response_model=List[Book]
)
def get_books():
    return books_db


@router.post("/",
    summary="Add a new book",
    description="Adds a new book to the digital library",
    response_model=Book
)
def add_book(book: Book):
    for b in books_db:
        if b.id == book.id:
            raise HTTPException(status_code=400, detail="Book with this id already exists")
    books_db.append(book)
    return book

@router.get("/{book_id}",
    summary="Get a specific book",
    description="Retrieves a specific book from the digital library by id",
    response_model=Book
)
def get_book(book_id: int):
    for book in books_db:
        if book.id == book_id:
            return book
    raise HTTPException(status_code=404, detail="Book not found")

@router.put(
    "/{book_id}",
    summary="Update a book",
    description="Updates the details of an existing book in the library",
    response_model=Book
)
def update_book(book_id: int, updated_book: Book):
    for index, book in enumerate(books_db):
        if book.id == book_id:
            books_db[index] = updated_book
            return updated_book
    raise HTTPException(status_code=404, detail="Book not found")

@router.delete(
    "/{book_id}",
    summary="Delete a book",
    description="Removes a book permanently from the digital library"
)
def delete_book(book_id: int):
    for book in books_db:
        if book.id == book_id:
            books_db.remove(book)
            return {"message": "Book deleted successfully"}
    raise HTTPException(status_code=404, detail="Book not found")