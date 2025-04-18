package serializer

import (
	"time"

	"github.com/iandanarko/concert/internal/model/concert"
)

type ConcertResponse struct {
	ID   uint64    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

// SerializeConcerts serialize response for get concerts list
func SerializeConcerts(datas []concert.Concert, total, offset, limit uint64) SuccessResponse[[]ConcertResponse] {
	result := make([]ConcertResponse, len(datas))
	for i, d := range datas {
		result[i] = ConcertResponse{
			ID:   d.ID,
			Name: d.Name,
			Date: d.Date,
		}
	}

	return SuccessResponse[[]ConcertResponse]{
		Data: result,
		Metadata: &ListMetadata{
			Total:  total,
			Offset: offset,
			Limit:  limit,
		},
	}
}
