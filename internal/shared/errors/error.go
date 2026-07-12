package errors

import "errors"

var (
	// Common
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("resource not found")
	ErrConflict       = errors.New("resource already exists")

	// Authentication
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")

	// Validation
	ErrValidation = errors.New("validation failed")
)
