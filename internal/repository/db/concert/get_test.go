package concertrepo_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/iandanarko/concert/internal/model/concert"
	"github.com/stretchr/testify/assert"
)

func TestConcert_GetAvailableConcerts(t *testing.T) {
	t.Run("Failed: error query", func(t *testing.T) {
		t.Parallel()
		expQuery := regexp.QuoteMeta("SELECT id, name, date, created_at, updated_at FROM concerts WHERE date > ? ORDER BY date ASC")
		suite := new(t)
		defer suite.db.Close()

		suite.mock.ExpectQuery(expQuery).WillReturnError(errors.New("unexpected error"))

		results, err := suite.repo.GetAvailableConcerts(context.TODO(), concert.GetAvailableSpec{Offset: 0, Limit: 10})
		assert.Error(t, err)
		assert.Empty(t, results)
	})

	t.Run("Success: success get with search", func(t *testing.T) {
		t.Parallel()
		expQuery := regexp.QuoteMeta("SELECT id, name, date, created_at, updated_at FROM concerts WHERE date > now() AND name LIKE (?%) ORDER BY date ASC")
		suite := new(t)
		defer suite.db.Close()

		suite.mock.ExpectQuery(expQuery).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "date", "created_at", "updated_at"}).
			AddRow(1, "Concert test", time.Now(), time.Now(), time.Now()))

		results, err := suite.repo.GetAvailableConcerts(context.TODO(), concert.GetAvailableSpec{Offset: 0, Limit: 10, Search: "Test"})
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})
}
