package booking

import (
	"context"

	cErr "github.com/iandanarko/concert/internal/error"
	"github.com/iandanarko/concert/internal/model/booking"
	"github.com/iandanarko/concert/internal/usecase"
	"github.com/labstack/gommon/log"
)

// BookingService handle business logic for booking concert tickets
type BookingService struct {
	bookingRepo usecase.BookRepo
	concertRepo usecase.GetConcertByIDRepo
}

// NewBookingService creates an instance for BookingService
func NewBookingService(
	bookingRepo usecase.BookRepo,
	concertRepo usecase.GetConcertByIDRepo,
) BookingService {
	return BookingService{
		bookingRepo: bookingRepo,
		concertRepo: concertRepo,
	}
}

// Book create a booking for concert
func (s BookingService) Book(ctx context.Context, spec booking.BookSpec) error {
	concert, err := s.concertRepo.GetByID(ctx, spec.ConcertID)
	if err != nil {
		log.Printf("[Book] Error find concert by id %d: %v", spec.ConcertID, err)
		return err
	}

	if concert == nil {
		return cErr.ErrConcertNotFound
	}

	err = s.bookingRepo.Book(ctx, spec)
	if err != nil {
		log.Printf("[Book] Error book concert %d: %v", spec.ConcertID, err)
		return err
	}

	return nil
}
