package apperror

import (
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrNotFound = errors.New("data not found")
)
