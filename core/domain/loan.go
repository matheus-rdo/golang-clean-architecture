package domain

import (
	"errors"
	"time"
)

// Loan loan entity
type Loan struct {
	ID            string    `json:"id"`
	BookID        string    `json:"bookId" binding:"required"`
	RequesterName string    `json:"requesterName" binding:"required"`
	ReturnDate    time.Time `json:"returnDate"`
}

var (
	// ErrExistingLoan books can't be borrowed twice
	ErrExistingLoan = errors.New("Sorry, this book already haves a pending loan")
)

// Validate validate loan
func (loan Loan) Validate() error {
	if loan.BookID == "" || loan.RequesterName == "" {
		return ErrInvalidEntity
	}

	return nil
}

// LoanUseCase usecase to manage loans
type LoanUseCase interface {
	HasLoan(bookID string) (*bool, error)
	BorrowBook(Loan) (*Loan, error)
	ReturnBook(bookID string) error
}

// LoanRepository repository to persist loan entity
type LoanRepository interface {
	Create(Loan) (*Loan, error)
	FindByBookID(bookID string) (*Loan, error)
	DeleteByBookID(bookID string) error
}
