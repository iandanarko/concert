package serializer

type SuccessResponse[T any] struct {
	Data     T             `json:"data"`
	Metadata *ListMetadata `json:"metadata"`
}

type ListMetadata struct {
	Total  uint64 `json:"total"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

type SuccessMessageResponse struct {
	Data MessageResponse
}

type MessageResponse struct {
	Message string `json:"message"`
}

func SerializeMessageResponse(message string) SuccessMessageResponse {
	return SuccessMessageResponse{
		Data: MessageResponse{Message: message},
	}
}
