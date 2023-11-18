package helper

import "errors"

// sentinel error
var (
	ErrUserNotFound = errors.New("error user not found")
	ErrDuplicateKey = errors.New("error duplicate key")
)

// custom error
var (
	// token
	ErrInvalidResetToken   = NewAppError(StatusUnauthorized, "error invalid reset token")
	ErrGenerateToken       = NewAppError(StatusInternalServerError, "error generating token")
	ErrInvalidTokenType    = NewAppError(StatusUnauthorized, "Invalid token type")
	ErrInvalidJWTToken     = NewAppError(StatusUnauthorized, "error invalid token")
	ErrAccessTokenExpired  = NewAppError(StatusUnauthorized, "AccessTokenExpired")
	ErrRefreshTokenExpired = NewAppError(StatusUnauthorized, "RefreshTokenExpired")
	ErrStepUpTokenExpired  = NewAppError(StatusUnauthorized, "StepUpTokenExpired")

	ErrInternalServer    = NewAppError(StatusInternalServerError, "internal server error")
	ErrInvalidAuthHeader = NewAppError(StatusBadRequest, "error invalid auth header")
	ErrInvalidBody       = NewAppError(StatusBadRequest, "error invalid body")

	ErrTransferSameUser = NewAppError(StatusBadRequest, "error transfer to same user")
	ErrUnauthorizedUser = NewAppError(StatusUnauthorized, "error unauthorized user")
	ErrCreatingNewUser  = NewAppError(StatusInternalServerError, "error creating new user")
	ErrDuplicateUser    = NewAppError(StatusBadRequest, "error email already exist")
	ErrCredential       = NewAppError(StatusBadRequest, "error incorrect email/password")

	ErrBoxNotFound = NewAppError(StatusBadRequest, "error box not found")
)
