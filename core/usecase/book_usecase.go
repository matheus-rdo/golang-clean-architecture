package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/matheushr97/golang-clean-architecture/core/domain"
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
	if err := book.Validate(); err != nil {
		return nil, err
	}

	now := time.Now()
	book.ID = uuid.New().String()
	book.CreatedAt = now
	book.UpdatedAt = now

	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

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
	book, err := uc.bookRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, domain.ErrNotFound
	}

	return book, nil
}

func (uc *bookUsecase) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()
	book, err := uc.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if book == nil {
		return domain.ErrNotFound
	}

	return uc.bookRepo.Delete(ctx, id)
}
