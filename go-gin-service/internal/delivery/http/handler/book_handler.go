package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vKousik/fastapi-digital-library/internal/domain"
	"github.com/vKousik/fastapi-digital-library/internal/usecase"
)

type BookHandler struct {
	usecase *usecase.BookUseCase
}

func NewBookHandler(usecase *usecase.BookUseCase) *BookHandler {
	return &BookHandler{usecase: usecase}
}
func (b *BookHandler) GetBooks(c *gin.Context) {
	books, err := b.usecase.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Books fetched successfully",
		"data":    books,
	})
}
func (b *BookHandler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	book, err := b.usecase.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book fetched successfully",
		"data":    book,
	})
}
func (b *BookHandler) CreateBook(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := b.usecase.CreateBook(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"data":    book,
	})
}
func (b *BookHandler) UpdateBook(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := b.usecase.UpdateBook(book); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"data":    book,
	})
}
func (b *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if err := b.usecase.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
		"data":    nil,
	})
}
