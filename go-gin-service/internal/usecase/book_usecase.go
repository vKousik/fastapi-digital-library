package usecase

import (
	"errors"

	"github.com/vKousik/fastapi-digital-library/internal/domain"
)

type BookUseCase struct {
	repo domain.BookRepository
}

func NewBookUseCase(repo domain.BookRepository) *BookUseCase {
	return &BookUseCase{repo: repo}
}

func (b *BookUseCase) GetBooks() ([]domain.Book, error) {
	return b.repo.GetBooks()
}

func (b *BookUseCase) GetBookById(id int) (domain.Book, error) {
	return b.repo.GetBookById(id)
}

func (b *BookUseCase) CreateBook(book domain.Book) error {
	if book.Title == "" {
		return errors.New("Title is required")
	}
	if book.Year < 1000 || book.Year > 2026 {
		return errors.New("Year must be between 1000 and 2026")
	}
	if len(book.ISBN) != 10 && len(book.ISBN) != 13 {
		return errors.New("ISBN format invalid")
	}
	return b.repo.CreateBook(book)
}

func (b *BookUseCase) UpdateBook(book domain.Book) error {
	return b.repo.UpdateBook(book)
}

func (b *BookUseCase) DeleteBook(id int) error {
	return b.repo.DeleteBook(id)
}
