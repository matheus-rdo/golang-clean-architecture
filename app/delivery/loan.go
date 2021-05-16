package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheushr97/golang-clean-architecture/app/handler/presenter"
	"github.com/matheushr97/golang-clean-architecture/core/domain"
)

// LoanHandler requests/responses handler
type LoanHandler struct {
	UseCase domain.LoanUseCase
}

// NewLoanHandler instantiates a new LoanHandler
func NewLoanHandler(router *gin.RouterGroup, loanUseCase domain.LoanUseCase) {
	handler := LoanHandler{
		UseCase: loanUseCase,
	}
	router.GET("/loan/books/:bookId", handler.hasLoan)
	router.POST("/loan/books", handler.borrowBook)
	router.DELETE("/loan/books/:bookId", handler.returnBook)
}

func (handler *LoanHandler) hasLoan(context *gin.Context) {
	hasLoan, err := handler.UseCase.HasLoan(context.Param("bookId"))
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}

	context.JSON(http.StatusOK, presenter.HasLoan{
		Exists: *hasLoan,
	})
}

func (handler *LoanHandler) borrowBook(context *gin.Context) {
	var loan domain.Loan
	if err := context.BindJSON(&loan); err != nil {
		return // throws bad request
	}

	newLoan, err := handler.UseCase.BorrowBook(loan)
	if err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}

	context.JSON(http.StatusCreated, newLoan)
}

func (handler *LoanHandler) returnBook(context *gin.Context) {
	if err := handler.UseCase.ReturnBook(context.Param("bookId")); err != nil {
		context.AbortWithStatusJSON(getStatusCode(err), buildResponseFromError(err))
		return
	}
	context.AbortWithStatus(http.StatusOK)
}
