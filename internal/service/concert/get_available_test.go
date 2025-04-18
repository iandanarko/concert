package concert_test

import (
	"context"
	"errors"
	"testing"

	"github.com/iandanarko/concert/internal/model/concert"
	svc "github.com/iandanarko/concert/internal/service/concert"
	mock_service "github.com/iandanarko/concert/test/mock/usecase"
	"github.com/stretchr/testify/assert"
)

type getavailableTest struct {
	repo    *mock_service.GetAvailableConcertListRepo
	service svc.GetAvailableService
}

func newGetavailableTest(t *testing.T) getavailableTest {
	repo := mock_service.NewGetAvailableConcertListRepo(t)
	service := svc.NewGetAvailable(repo)
	return getavailableTest{
		repo:    repo,
		service: service,
	}
}

func TestGetAvailableService_GetAvailable(t *testing.T) {
	t.Run("Failed: error on repository", func(t *testing.T) {
		t.Parallel()
		ctx := context.TODO()
		spec := concert.GetAvailableSpec{}
		expErr := errors.New("some error")
		suite := newGetavailableTest(t)
		suite.repo.EXPECT().GetAvailableConcerts(ctx, spec).Return([]concert.Concert{}, expErr).Once()
		result, err := suite.service.GetAvailable(ctx, spec)
		assert.ErrorIs(t, expErr, err)
		assert.Empty(t, result)
	})

	t.Run("Success: success get", func(t *testing.T) {
		t.Parallel()
		ctx := context.TODO()
		spec := concert.GetAvailableSpec{}
		exp := []concert.Concert{{ID: 1}}
		suite := newGetavailableTest(t)
		suite.repo.EXPECT().GetAvailableConcerts(ctx, spec).Return(exp, nil).Once()
		result, err := suite.service.GetAvailable(ctx, spec)
		assert.NoError(t, err)
		assert.Equal(t, exp, result)
	})
}
