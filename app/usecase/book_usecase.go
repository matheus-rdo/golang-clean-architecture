package usecase

import (
	"context"
	"time"

	"github.com/matheushr97/golang-clean-architecture/domain"
)

type bookUsecase struct {
	bookRepo       domain.BookRepository
	contextTimeout time.Duration
}

// NewBookUseCase impl of domain.BookUseCase
func NewBookUseCase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookUseCase {
	return &bookUsecase{
		bookRepo:       bookRepository,
		contextTimeout: timeout,
	}
}

func (uc *bookUsecase) Create(book domain.Book) (res *domain.Book, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	// ALL BUSINESS RULES HERE

	return uc.bookRepo.Create(ctx, book)
}

func (uc *bookUsecase) Fetch() (res *[]domain.Book, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	return uc.bookRepo.Fetch(ctx)
}

func (uc *bookUsecase) GetByID(id string) (*domain.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	return uc.bookRepo.GetByID(ctx, id)
}
