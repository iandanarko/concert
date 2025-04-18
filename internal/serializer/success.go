package serializer

type SuccessResponse[T any] struct {
	Data     T   `json:"data"`
	Metadata any `json:"metadata"`
}

type ListMetadata struct {
	Total  uint64 `json:"total"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}
