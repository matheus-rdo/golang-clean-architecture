package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheushr97/golang-clean-architecture/app/handler/presenter"
	"github.com/matheushr97/golang-clean-architecture/domain"
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
	router.GET("/books", handler.FetchBooks)
	router.POST("/books", handler.AddBook)
}

// FetchBooks will fetch all books
func (handler *BookHandler) FetchBooks(context *gin.Context) {
	books, err := handler.BookUseCase.Fetch()
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}
	dtos := parseToDTOs(*books)
	context.JSON(http.StatusOK, dtos)
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
	dto := parseToDTO(*created)
	context.SecureJSON(http.StatusCreated, dto)
}

// GetBookByID finds a book with given ID
func (handler *BookHandler) GetBookByID(response http.ResponseWriter, request *http.Request) {
	//TODO: Impl
}

func parseToDTOs(books []domain.Book) []presenter.Book {
	dtos := make([]presenter.Book, len(books))
	for _, book := range books {
		dtos = append(dtos, parseToDTO(book))
	}
	return dtos
}

func parseToDTO(book domain.Book) presenter.Book {
	return presenter.Book{
		ID:      book.ID,
		Title:   book.Title,
		Content: book.Content,
		Author:  book.Author,
	}
}
