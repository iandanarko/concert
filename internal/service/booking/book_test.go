package booking_test

import (
	"context"
	"errors"
	"testing"

	"github.com/iandanarko/concert/internal/model/booking"
	"github.com/iandanarko/concert/internal/model/concert"
	svc "github.com/iandanarko/concert/internal/service/booking"
	mock_service "github.com/iandanarko/concert/test/mock/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type bookingTest struct {
	bookingRepo *mock_service.BookRepo
	concertRepo *mock_service.GetConcertByIDRepo
	service     svc.BookingService
}

func newBookingTest(t *testing.T) bookingTest {
	bookingRepo := mock_service.NewBookRepo(t)
	concertRepo := mock_service.NewGetConcertByIDRepo(t)
	service := svc.NewBookingService(bookingRepo, concertRepo)
	return bookingTest{
		bookingRepo: bookingRepo,
		concertRepo: concertRepo,
		service:     service,
	}
}

func TestBookingService_Book(t *testing.T) {
	t.Run("Failed: error find concert", func(t *testing.T) {
		t.Parallel()
		suite := newBookingTest(t)
		suite.concertRepo.EXPECT().GetByID(mock.Anything, mock.Anything).Return(nil, errors.New("some error")).Once()
		err := suite.service.Book(context.TODO(), booking.BookSpec{})
		assert.Error(t, err)
	})

	t.Run("Failed: concert not found", func(t *testing.T) {
		t.Parallel()
		suite := newBookingTest(t)
		suite.concertRepo.EXPECT().GetByID(mock.Anything, mock.Anything).Return(nil, nil).Once()
		err := suite.service.Book(context.TODO(), booking.BookSpec{})
		assert.Error(t, err)
	})

	t.Run("Failed: error book", func(t *testing.T) {
		t.Parallel()
		suite := newBookingTest(t)
		suite.concertRepo.EXPECT().GetByID(mock.Anything, mock.Anything).Return(&concert.Concert{}, nil).Once()
		suite.bookingRepo.EXPECT().Book(mock.Anything, mock.Anything).Return(errors.New("some error")).Once()
		err := suite.service.Book(context.TODO(), booking.BookSpec{})
		assert.Error(t, err)
	})

	t.Run("Success: success book", func(t *testing.T) {
		t.Parallel()
		suite := newBookingTest(t)
		suite.concertRepo.EXPECT().GetByID(mock.Anything, mock.Anything).Return(&concert.Concert{}, nil).Once()
		suite.bookingRepo.EXPECT().Book(mock.Anything, mock.Anything).Return(nil).Once()
		err := suite.service.Book(context.TODO(), booking.BookSpec{})
		assert.NoError(t, err)
	})
}
