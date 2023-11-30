package errmsg

import "errors"

var (
	ErrUserNotFound    = errors.New("error user not found")
	ErrDuplicateKey    = errors.New("error duplicate key")
	ErrUserVerifyEmail = errors.New("error user not found / already verified")
)
