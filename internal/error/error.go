package error

import "errors"

// CustomError used for custom error
type CustomError error

var (
	// ErrUnauthorized is returned when error is unauthorized
	ErrUnauthorized CustomError = errors.New("Unauthorized")
	// ErrBadRequest is returned when error bad request
	ErrBadRequest CustomError = errors.New("Bad Request")
	// ErrInternalServer is returned when error is unknown
	ErrInternalServer CustomError = errors.New("Internal Server")
)
