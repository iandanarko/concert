package serializer

// ErrorResponse used to hold response for error
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse serialize new ErrorResponse
func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Error: err.Error(),
	}
}
