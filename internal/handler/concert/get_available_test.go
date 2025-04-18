package concert_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iandanarko/concert/internal/handler/concert"
	model "github.com/iandanarko/concert/internal/model/concert"
	mock_service "github.com/iandanarko/concert/test/mock/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetavailableHandler_Handle(t *testing.T) {
	t.Run("Failed: error bind query", func(t *testing.T) {
		t.Parallel()
		url := "/concerts?offset=satu"
		svc := mock_service.NewGetAvailableConcertListUseCase(t)
		handler := concert.NewGetAvailableHandler(svc)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, rec)
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Failed: error service", func(t *testing.T) {
		t.Parallel()
		url := "/concerts?offset=0"
		svc := mock_service.NewGetAvailableConcertListUseCase(t)
		handler := concert.NewGetAvailableHandler(svc)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		svc.EXPECT().GetAvailable(mock.Anything, mock.Anything).Return([]model.Concert{}, 0, errors.New("some error")).Once()

		e := echo.New()
		ctx := e.NewContext(req, rec)
		err := handler.Handle(ctx)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("Failed: error service", func(t *testing.T) {
		t.Parallel()
		url := "/concerts?offset=0&limit=1&search=test"
		svc := mock_service.NewGetAvailableConcertListUseCase(t)
		handler := concert.NewGetAvailableHandler(svc)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		svc.EXPECT().GetAvailable(mock.Anything, mock.Anything).Return([]model.Concert{{ID: 1}}, 1, nil).Once()

		e := echo.New()
		ctx := e.NewContext(req, rec)
		err := handler.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
