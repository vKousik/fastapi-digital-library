package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vKousik/fastapi-digital-library/internal/delivery/http/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.BookHandler) {
	books := r.Group("/books")
	{
		books.POST("", h.CreateBook)
		books.GET("", h.GetBooks)
		books.GET("/:id", h.GetBookById)
		books.PUT("", h.UpdateBook)
		books.DELETE("/:id", h.DeleteBook)
	}
}
