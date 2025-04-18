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
	// ErrConcertNotFound returned when concert not found
	ErrConcertNotFound CustomError = errors.New("concert not found")
	// ErrOutOfTickets returned when tickets out of stock
	ErrOutOfTickets CustomError = errors.New("Out Of Tickets")
	// ErrTicketWindowNotFound returned when ticket window not found
	ErrTicketWindowNotFound CustomError = errors.New("Ticket Window not available")
)
