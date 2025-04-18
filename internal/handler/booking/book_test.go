package booking_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	cErr "github.com/iandanarko/concert/internal/error"
	"github.com/iandanarko/concert/internal/handler/booking"
	"github.com/iandanarko/concert/internal/middleware"
	"github.com/iandanarko/concert/internal/model"
	mock_service "github.com/iandanarko/concert/test/mock/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookingHandler_Book(t *testing.T) {
	url := "/bookings"
	actor := &model.Actor{UserID: 1}
	t.Run("Failed: error bind body", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal("invalid")

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Failed: error invalid concert id", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal(booking.BookBodyRequest{})

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Failed: error invalid quantity", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal(booking.BookBodyRequest{ConcertID: 1})

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Failed: error service", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal(booking.BookBodyRequest{ConcertID: 1, Quantity: 1})

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		svc.EXPECT().Book(mock.Anything, mock.Anything).Return(errors.New("some error")).Once()
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Failed: error concert not found", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal(booking.BookBodyRequest{ConcertID: 1, Quantity: 1})

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		svc.EXPECT().Book(mock.Anything, mock.Anything).Return(cErr.ErrConcertNotFound).Once()
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Failed: error concert not found", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal(booking.BookBodyRequest{ConcertID: 1, Quantity: 1})

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		svc.EXPECT().Book(mock.Anything, mock.Anything).Return(cErr.ErrOutOfTickets).Once()
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("Success: success", func(t *testing.T) {
		t.Parallel()
		svc := mock_service.NewBookUseCase(t)
		handler := booking.NewBookingHandler(svc)
		body, _ := json.Marshal(booking.BookBodyRequest{ConcertID: 1, Quantity: 1})

		req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.Set(middleware.Actor, actor)
		svc.EXPECT().Book(mock.Anything, mock.Anything).Return(nil).Once()
		err := handler.Handle(ctx)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
