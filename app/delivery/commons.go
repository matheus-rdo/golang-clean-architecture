package handler

import (
	"net/http"

	"github.com/matheushr97/golang-clean-architecture/core/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrInvalidEntity:
		return http.StatusUnprocessableEntity
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func buildResponseFromError(err error) ResponseError {
	return ResponseError{
		Message: err.Error(),
	}
}
