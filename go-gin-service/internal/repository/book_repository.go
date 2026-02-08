package repository

import (
	"errors"

	"github.com/vKousik/fastapi-digital-library/internal/domain"
)

type InMemoryBookRepository struct {
	books []domain.Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: []domain.Book{},
	}
}
func (r *InMemoryBookRepository) CreateBook(book domain.Book) error {
	for _, b := range r.books {
		if b.ID == book.ID {
			return errors.New("Duplicate Id")
		}
	}
	r.books = append(r.books, book)
	return nil
}
func (r *InMemoryBookRepository) GetBooks() ([]domain.Book, error) {
	return r.books, nil
}
func (r *InMemoryBookRepository) GetBookById(id int) (domain.Book, error) {
	for _, b := range r.books {
		if id == b.ID {
			return b, nil
		}
	}
	return domain.Book{}, errors.New("Book Not found")
}

func (r *InMemoryBookRepository) UpdateBook(book domain.Book) error {
	for i, b := range r.books {
		if b.ID == book.ID {
			r.books[i] = book
			return nil
		}
	}
	return errors.New("Not found")
}

func (r *InMemoryBookRepository) DeleteBook(id int) error {
	for i, b := range r.books {
		if b.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}
