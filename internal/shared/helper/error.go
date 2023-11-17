package helper

import (
	"fmt"
)

type (
	AppError struct {
		Code    AppErrorCode `json:"code"`
		Message string       `json:"message"`
	}

	AppErrorCode int

	ErrorDto struct {
		Message string `json:"message"`
	}
)

const (
	StatusBadRequest AppErrorCode = iota
	StatusForbidden
	StatusUnauthorized
	StatusInternalServerError
)

func (ce *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.Code, ce.Message)
}

func (ce *AppError) ToErrorDto() ErrorDto {
	return ErrorDto{
		Message: ce.Message,
	}
}

func NewAppError(code AppErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
