package booking

import (
	"errors"
	"net/http"

	cErr "github.com/iandanarko/concert/internal/error"
	"github.com/iandanarko/concert/internal/middleware"
	"github.com/iandanarko/concert/internal/model"
	"github.com/iandanarko/concert/internal/model/booking"
	"github.com/iandanarko/concert/internal/serializer"
	"github.com/iandanarko/concert/internal/usecase"
	"github.com/labstack/echo/v4"
)

type bookingHandler struct {
	svc usecase.BookUseCase
}

// NewBookingHandler creates an instance of bookingHandler
func NewBookingHandler(svc usecase.BookUseCase) bookingHandler {
	return bookingHandler{
		svc: svc,
	}
}

// BookTickets godoc
//
//	@Security		JWTAuth
//	@Summary		book tickets
//	@Description book tickets.
//	@Tags			booking
//	@Accept			json
//	@Produce		json
//	@Param body body BookBodyRequest true "body request"
//	@Success		200	{object} serializer.SuccessMessageResponse	"success booking"
//	@Router			/bookings [post]
func (h bookingHandler) Handle(echoCtx echo.Context) error {
	actor := echoCtx.Get(middleware.Actor).(*model.Actor)
	ctx := echoCtx.Request().Context()

	body := BookBodyRequest{}
	if err := echoCtx.Bind(&body); err != nil {
		echoCtx.JSON(http.StatusBadRequest, serializer.NewErrorResponse(cErr.ErrBadRequest))
		return err
	}

	if err := body.validate(); err != nil {
		echoCtx.JSON(http.StatusBadRequest, serializer.NewErrorResponse(err))
		return err
	}

	spec := booking.BookSpec{
		ConcertID:  body.ConcertID,
		UserID:     actor.UserID,
		NumTickets: body.Quantity,
	}

	err := h.svc.Book(ctx, spec)
	if err != nil {
		if err == cErr.ErrConcertNotFound {
			echoCtx.JSON(http.StatusNotFound, serializer.NewErrorResponse(err))
		} else if err == cErr.ErrOutOfTickets || err == cErr.ErrTicketWindowNotFound {
			echoCtx.JSON(http.StatusUnprocessableEntity, serializer.NewErrorResponse(err))
		} else {
			echoCtx.JSON(http.StatusInternalServerError, serializer.NewErrorResponse(cErr.ErrInternalServer))
		}

		return err
	}

	echoCtx.JSON(http.StatusOK, serializer.SerializeMessageResponse("Success"))
	return nil
}

// BookBodyRequest used to hold data for booking request
type BookBodyRequest struct {
	ConcertID uint64 `json:"concert_id" validate:"required"`
	Quantity  uint64 `json:"quantity" validate:"required"`
}

func (r BookBodyRequest) validate() error {
	if r.ConcertID == 0 {
		return errors.New("invalid concert id")
	}

	if r.Quantity == 0 {
		return errors.New("invalid quantity")
	}
	return nil
}
