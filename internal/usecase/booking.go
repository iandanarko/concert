package usecase

import (
	"context"

	"github.com/iandanarko/concert/internal/model/booking"
)

type BookRepo interface {
	Book(ctx context.Context, spec booking.BookSpec) error
}

type BookUseCase interface {
	Book(ctx context.Context, spec booking.BookSpec) error
}
