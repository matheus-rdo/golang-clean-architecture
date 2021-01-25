package domain

import (
	"context"
	"time"
)

// Book book entity
type Book struct {
	ID        string    `json:"bookId"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Author    string    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate validate book
func (book Book) Validate() error {
	if book.Title == "" || book.Author == "" {
		return ErrInvalidEntity
	}

	return nil
}

// BookUseCase usecases to manage book entity
type BookUseCase interface {
	Create(Book) (res *Book, err error)
	Fetch() (res *[]Book, err error)
	GetByID(id string) (*Book, error)
	Delete(id string) error
}

// BookRepository repository to save book entity
type BookRepository interface {
	Create(ctx context.Context, book Book) (res *Book, err error)
	Fetch(ctx context.Context) (res *[]Book, err error)
	GetByID(ctx context.Context, id string) (*Book, error)
	Delete(ctx context.Context, id string) error
}
