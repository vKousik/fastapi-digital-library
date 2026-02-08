package domain

type BookRepository interface {
	GetBookById(id int) (Book, error)
	GetBooks() ([]Book, error)
	CreateBook(book Book) error
	UpdateBook(book Book) error
	DeleteBook(id int) error
}
