package helper

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound = errors.New("error user not found")
)

var (
	ErrInternalServer    = NewAppError(StatusInternalServerError, "internal server error")
	ErrInvalidResetToken = NewAppError(StatusUnauthorized, "error invalid reset token")
	ErrGenerateToken     = NewAppError(StatusInternalServerError, "error generating token")
	ErrInvalidJWTToken   = NewAppError(StatusUnauthorized, "error invalid token")
	ErrInvalidAuthHeader = NewAppError(StatusBadRequest, "error invalid auth header")
	ErrInvalidBody       = NewAppError(StatusBadRequest, "error invalid body")

	ErrTransferSameUser = NewAppError(StatusBadRequest, "error transfer to same user")
	ErrUnauthorizedUser = NewAppError(StatusUnauthorized, "error unauthorized user")
	ErrCreatingNewUser  = NewAppError(StatusInternalServerError, "error creating new user")
	ErrDuplicateUser    = NewAppError(StatusBadRequest, "error email already exist")
	ErrCredential       = NewAppError(StatusBadRequest, "error incorrect email/password")

	ErrBoxNotFound = NewAppError(StatusBadRequest, "error box not found")
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
