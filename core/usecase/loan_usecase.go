package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/matheushr97/golang-clean-architecture/core/domain"
)

type loanUsecase struct {
	bookRepository domain.BookRepository
	loanRepository domain.LoanRepository
	contextTimeout time.Duration
}

// NewLoanUseCase impl of domain.LoanUseCase
func NewLoanUseCase(bookRepository domain.BookRepository, loanRepository domain.LoanRepository) domain.LoanUseCase {
	return &loanUsecase{
		bookRepository: bookRepository,
		loanRepository: loanRepository,
		contextTimeout: time.Second * 3,
	}
}

func (uc loanUsecase) HasLoan(bookID string) (*bool, error) {
	if _, err := uc.requireBook(bookID); err != nil {
		return nil, err
	}

	loan, err := uc.loanRepository.FindByBookID(bookID)
	if err != nil {
		return nil, err
	}
	hasLoan := loan != nil
	return &hasLoan, nil
}

func (uc loanUsecase) BorrowBook(loan domain.Loan) (*domain.Loan, error) {
	if err := loan.Validate(); err != nil {
		return nil, err
	}

	// Cannot borrow a book that have a pending loan
	hasLoan, err := uc.HasLoan(loan.BookID)
	if err != nil {
		return nil, err
	}
	if *hasLoan {
		return nil, domain.ErrExistingLoan
	}
	loan.ID = uuid.New().String()
	return uc.loanRepository.Create(loan)
}

func (uc loanUsecase) ReturnBook(bookID string) error {
	if _, err := uc.requireBook(bookID); err != nil {
		return err
	}
	return uc.loanRepository.DeleteByBookID(bookID)
}

func (uc loanUsecase) requireBook(bookID string) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	book, err := uc.bookRepository.GetByID(ctx, bookID)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, domain.ErrNotFound
	}
	return book, nil
}
