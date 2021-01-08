package handler

import (
	"encoding/json"
	"net/http"

	"github.com/matheushr97/golang-clean-architecture/app/handler/presenter"
	"github.com/matheushr97/golang-clean-architecture/domain"
	"github.com/matheushr97/golang-clean-architecture/infra/web"
)

// BookHandler requests/responses handler
type BookHandler struct {
	BookUseCase domain.BookUseCase
}

// NewBookHandler instantiates a new BookHandler
func NewBookHandler(router web.Router, bookUseCase domain.BookUseCase) {
	handler := BookHandler{
		BookUseCase: bookUseCase,
	}
	router.GET("/api/books", handler.FetchBooks)
	//router.GET("/books/:id", handler.GetBookByID)
	router.POST("/api/books", handler.AddBook)
}

// FetchBooks will fetch all books
func (handler *BookHandler) FetchBooks(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	books, err := handler.BookUseCase.Fetch()
	if err != nil {
		response.WriteHeader(getStatusCode(err))
		json.NewEncoder(response).Encode(buildResponseFromError(err))
		return
	}
	dtos := parseToDTOs(*books)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(dtos)
}

// GetBookByID finds a book with given ID
func (handler *BookHandler) GetBookByID(response http.ResponseWriter, request *http.Request) {
	//TODO: Impl - does not have get url param
}

// AddBook handles a add book request
func (handler *BookHandler) AddBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var book domain.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		response.WriteHeader(getStatusCode(err))
		json.NewEncoder(response).Encode(buildResponseFromError(err))
		return
	}
	created, err := handler.BookUseCase.Create(book)
	if err != nil {
		response.WriteHeader(getStatusCode(err))
		json.NewEncoder(response).Encode(buildResponseFromError(err))
		return
	}
	bookDTO := presenter.Book{
		ID:      created.ID,
		Title:   created.Title,
		Content: created.Content,
		Author:  created.Author,
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(bookDTO)

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
