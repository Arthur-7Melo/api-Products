package config

import (
	"net/http"
)

type ProductError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Cause   []Causes `json:"cause,omitempty"`
}

type Causes struct{
	Field string `json:"field"`
	Message string `json:"message"`
}

func NewProductError(message string, code int, cause []Causes) *ProductError {
	return &ProductError{
		Message: message,
		Code:    code,
		Cause:   cause,
	}
}

func (pe *ProductError) Error() string {
	return pe.Message
}

func NewBadRequestError(message string) *ProductError{
	return &ProductError{
		Message: message,
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, cause []Causes) *ProductError {
	return &ProductError{
		Message: message,
		Code:    http.StatusBadRequest,
		Cause: cause,
	}
}

func NewInternalServerError(message string) *ProductError {
	return &ProductError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ProductError{
	return &ProductError{
		Message: message,
		Code: http.StatusNotFound,
	}
}

