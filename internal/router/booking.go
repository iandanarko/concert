package router

import (
	"database/sql"
	"net/http"

	"github.com/iandanarko/concert/internal/handler/booking"
	bookingrepo "github.com/iandanarko/concert/internal/repository/db/booking"
	concertrepo "github.com/iandanarko/concert/internal/repository/db/concert"
	bookingsvc "github.com/iandanarko/concert/internal/service/booking"
)

func bookingRouter(db *sql.DB) []route {
	concertRepo := concertrepo.New(db)
	bookingRepo := bookingrepo.New(db)
	bookingSvc := bookingsvc.NewBookingService(bookingRepo, concertRepo)
	bookingHandler := booking.NewBookingHandler(bookingSvc)
	return []route{
		{
			method:  http.MethodPost,
			path:    "/bookings",
			handler: bookingHandler.Handle,
		},
	}
}
