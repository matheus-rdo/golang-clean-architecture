package main

import (
	"time"

	"github.com/matheushr97/golang-clean-architecture/app"
	"github.com/matheushr97/golang-clean-architecture/app/handler"
	"github.com/matheushr97/golang-clean-architecture/app/repository"
	"github.com/matheushr97/golang-clean-architecture/app/usecase"
	"github.com/matheushr97/golang-clean-architecture/infra/web"
)

func main() {
	// Dependencies
	bookRepository := repository.NewBookRepository()
	bookUseCase := usecase.NewBookUseCase(bookRepository, time.Second*3)

	// Application routes
	router := web.NewGinRouter()
	handler.NewBookHandler(router, bookUseCase)

	router.SERVE(":" + app.ENV.APIPort)
}
