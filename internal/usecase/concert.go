package usecase

import (
	"context"

	"github.com/iandanarko/concert/internal/model/concert"
)

type GetAvailableConcertListRepo interface {
	GetAvailableConcerts(ctx context.Context, spec concert.GetAvailableSpec) ([]concert.Concert, error)
}

type GetAvailableConcertListUseCase interface {
}
