package concert

import (
	"context"
	"log"

	"github.com/iandanarko/concert/internal/model/concert"
	"github.com/iandanarko/concert/internal/usecase"
)

// GetAvailable is service to handle business logic for get avaiable concert list
type GetAvailableService struct {
	repo usecase.GetAvailableConcertListRepo
}

// NewGetAvailable creates new instance of GetAvailableService
func NewGetAvailable(
	repo usecase.GetAvailableConcertListRepo,
) GetAvailableService {
	return GetAvailableService{
		repo: repo,
	}
}

// GetAvailable list available concerts
func (s GetAvailableService) GetAvailable(ctx context.Context, spec concert.GetAvailableSpec) ([]concert.Concert, uint64, error) {
	results, total, err := s.repo.GetAvailableConcerts(ctx, spec)
	if err != nil {
		log.Printf("[GetAvailable] Error get available concerts: %v", err)
	}

	return results, total, err
}
