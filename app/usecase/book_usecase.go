package usecase

import (
	"time"

	"github.com/matheushr97/golang-clean-architecture/domain"
)

type bookUsecase struct {
	bookRepo       domain.BookRepository
	contextTimeout time.Duration
}
