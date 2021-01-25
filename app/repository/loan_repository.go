package repository

import (
	"github.com/matheushr97/golang-clean-architecture/core/domain"
)

type loanInMemoryRepository struct {
	bookLoansMap map[string]*domain.Loan
}

// NewLoanInMemoryRepository creates a new loan in memory repository
func NewLoanInMemoryRepository() domain.LoanRepository {
	return &loanInMemoryRepository{}
}

func (repository *loanInMemoryRepository) Create(loan domain.Loan) (*domain.Loan, error) {
	repository.bookLoansMap[loan.BookID] = &loan
	return &loan, nil
}

func (repository *loanInMemoryRepository) FindByBookID(bookID string) (*domain.Loan, error) {
	return repository.bookLoansMap[bookID], nil
}

func (repository *loanInMemoryRepository) DeleteByBookID(bookID string) error {
	repository.bookLoansMap[bookID] = nil
	return nil
}
