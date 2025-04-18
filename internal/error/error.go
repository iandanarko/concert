package error

import "errors"

// CustomError used for custom error
type CustomError error

var (
	// ErrUnauthorized is returned when error is unauthorized
	ErrUnauthorized CustomError = errors.New("Unauthorized")
)
