package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheushr97/golang-clean-architecture/app/handler/presenter"
	"github.com/matheushr97/golang-clean-architecture/core/domain"
)

// BookHandler requests/responses handler
type BookHandler struct {
	BookUseCase domain.BookUseCase
}

// NewBookHandler instantiates a new BookHandler
func NewBookHandler(router *gin.RouterGroup, bookUseCase domain.BookUseCase) {
	handler := BookHandler{
		BookUseCase: bookUseCase,
	}
	router.POST("/books", handler.AddBook)
	router.GET("/books", handler.FetchBooks)
	router.GET("/books/:id", handler.FetchBookByID)
	router.DELETE("/books/:id", handler.DeleteByID)
}

// FetchBooks will fetch all books
func (handler *BookHandler) FetchBooks(context *gin.Context) {
	books, err := handler.BookUseCase.Fetch()
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}
	context.JSON(http.StatusOK, parseToDTOs(books))
}

// FetchBookByID finds a book with given ID
func (handler *BookHandler) FetchBookByID(context *gin.Context) {
	book, err := handler.BookUseCase.GetByID(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}
	context.JSON(http.StatusOK, parseToDTO(book))
}

// DeleteByID deletes a book with given ID
func (handler *BookHandler) DeleteByID(context *gin.Context) {
	err := handler.BookUseCase.Delete(context.Param("id"))
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}
	context.JSON(http.StatusOK, nil)
}

// AddBook handles a add book request
func (handler *BookHandler) AddBook(context *gin.Context) {
	var entity domain.Book
	if err := context.BindJSON(&entity); err != nil {
		return // throws bad request
	}
	created, err := handler.BookUseCase.Create(entity)
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}
	context.SecureJSON(http.StatusCreated, parseToDTO(created))
}

func parseToDTOs(books *[]domain.Book) []presenter.Book {
	dtos := make([]presenter.Book, len(*books))
	for i, book := range *books {
		dtos[i] = parseToDTO(&book)
	}
	return dtos
}

func parseToDTO(book *domain.Book) presenter.Book {
	return presenter.Book{
		ID:      book.ID,
		Title:   book.Title,
		Content: book.Content,
		Author:  book.Author,
	}
}
