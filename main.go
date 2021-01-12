package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matheushr97/golang-clean-architecture/app"
	"github.com/matheushr97/golang-clean-architecture/app/handler"
	"github.com/matheushr97/golang-clean-architecture/app/repository"
	"github.com/matheushr97/golang-clean-architecture/app/usecase"
)

func main() {
	// Dependencies
	bookRepository := repository.NewBookRepository()
	bookUseCase := usecase.NewBookUseCase(bookRepository, time.Second*3)

	// Application routes
	engine := gin.New()
	apiRouter := engine.Group("/api")
	handler.NewBookHandler(apiRouter, bookUseCase)

	engine.Run(":" + app.ENV.APIPort)
}
