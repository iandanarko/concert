package concert

import (
	"net/http"

	cErr "github.com/iandanarko/concert/internal/error"
	"github.com/iandanarko/concert/internal/model/concert"
	"github.com/iandanarko/concert/internal/serializer"
	"github.com/iandanarko/concert/internal/usecase"
	"github.com/labstack/echo/v4"
)

type getavailableHandler struct {
	svc usecase.GetAvailableConcertListUseCase
}

// NewGetAvailableHandler creates an instance of getavailableHandler
func NewGetAvailableHandler(svc usecase.GetAvailableConcertListUseCase) getavailableHandler {
	return getavailableHandler{
		svc: svc,
	}
}

// GetAvailableQuery hold query params data for get available list
type GetAvailableQuery struct {
	Search string `query:"search"`
	Offset uint64 `query:"offset"`
	Limit  uint64 `query:"limt"`
}

// GetAvailableConcerts godoc
//
//	@Security		JWTAuth
//	@Summary		Get Available concerts list
//	@Description	Retrieve available concerts list.
//	@Tags			concert
//	@Accept			json
//	@Produce		json
//	@Param search query string false "search"
//	@Param offset query int false "offset"
//	@Param limit query int false "limit"
//	@Success		200	{object} serializer.SuccessResponse[serializer.ConcertResponse]	"retrieve concerts successfully"
//	@Router			/concerts [get]
func (h getavailableHandler) Handle(echoCtx echo.Context) error {
	ctx := echoCtx.Request().Context()
	q := GetAvailableQuery{}
	if err := echoCtx.Bind(&q); err != nil {
		echoCtx.JSON(http.StatusBadRequest, serializer.NewErrorResponse(cErr.ErrBadRequest))
		return err
	}

	spec := concert.GetAvailableSpec{Search: q.Search, Offset: q.Offset, Limit: q.Limit}
	results, total, err := h.svc.GetAvailable(ctx, spec)
	if err != nil {
		echoCtx.JSON(http.StatusInternalServerError, serializer.NewErrorResponse(cErr.ErrInternalServer))
		return err
	}

	echoCtx.JSON(http.StatusOK, serializer.SerializeConcerts(results, total, spec.Offset, spec.GetLimit()))
	return nil
}
