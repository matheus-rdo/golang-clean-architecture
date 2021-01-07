package domain

import (
	"context"
	"time"
)

// Book book entity
type Book struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    string    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// BookUseCase repository to save book entity
type BookUseCase interface {
	Create(Book) (res *Book, err error)
	Fetch() (res *[]Book, err error)
	GetByID(id string) (*Book, error)
}

// BookRepository repository to save book entity
type BookRepository interface {
	Create(ctx context.Context, book Book) (res *Book, err error)
	Fetch(ctx context.Context) (res *[]Book, err error)
	GetByID(ctx context.Context, id string) (*Book, error)
}
