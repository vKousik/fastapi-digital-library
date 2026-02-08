package main

import (
	"github.com/gin-gonic/gin"
	httpDelivery "github.com/vKousik/fastapi-digital-library/internal/delivery/http"
	"github.com/vKousik/fastapi-digital-library/internal/delivery/http/handler"
	"github.com/vKousik/fastapi-digital-library/internal/delivery/http/middleware"
	"github.com/vKousik/fastapi-digital-library/internal/repository"
	"github.com/vKousik/fastapi-digital-library/internal/usecase"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	repo := repository.NewInMemoryBookRepository()
	usecase := usecase.NewBookUseCase(repo)
	handler := handler.NewBookHandler(usecase)
	httpDelivery.RegisterRoutes(r, handler)

	r.Run(":8080")
}
