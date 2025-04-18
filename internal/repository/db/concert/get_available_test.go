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
		expQuery := regexp.QuoteMeta("SELECT id, name, date, created_at, updated_at FROM concerts WHERE date > ? ORDER BY date ASC ORDER BY date ASC LIMIT 10 OFFSET 0")
		suite := new(t)
		defer suite.db.Close()

		suite.mock.ExpectQuery(expQuery).WillReturnError(errors.New("unexpected error"))

		results, total, err := suite.repo.GetAvailableConcerts(context.TODO(), concert.GetAvailableSpec{Offset: 0, Limit: 10})
		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Zero(t, total)
	})

	t.Run("Failed: error count", func(t *testing.T) {
		t.Parallel()
		expQuery := regexp.QuoteMeta("SELECT id, name, date, created_at, updated_at FROM concerts WHERE date > now() AND name LIKE ? ORDER BY date ASC LIMIT 10 OFFSET 0")
		expCount := regexp.QuoteMeta("SELECT COUNT(*) FROM concerts WHERE date > now() AND name LIKE ?")
		suite := new(t)
		defer suite.db.Close()

		suite.mock.ExpectQuery(expQuery).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "date", "created_at", "updated_at"}).
			AddRow(1, "Concert test", time.Now(), time.Now(), time.Now()))
		suite.mock.ExpectQuery(expCount).WillReturnError(errors.New("some error"))

		results, total, err := suite.repo.GetAvailableConcerts(context.TODO(), concert.GetAvailableSpec{Offset: 0, Limit: 10, Search: "Test"})
		assert.Error(t, err)
		assert.Empty(t, results)
		assert.Zero(t, total)
	})

	t.Run("Success: success get with search", func(t *testing.T) {
		t.Parallel()
		expQuery := regexp.QuoteMeta("SELECT id, name, date, created_at, updated_at FROM concerts WHERE date > now() AND name LIKE ? ORDER BY date ASC LIMIT 10 OFFSET 0")
		expCount := regexp.QuoteMeta("SELECT COUNT(*) FROM concerts WHERE date > now() AND name LIKE ?")
		suite := new(t)
		defer suite.db.Close()
		expTotal := uint64(1)

		suite.mock.ExpectQuery(expQuery).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "date", "created_at", "updated_at"}).
			AddRow(1, "Concert test", time.Now(), time.Now(), time.Now()))
		suite.mock.ExpectQuery(expCount).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expTotal))

		results, total, err := suite.repo.GetAvailableConcerts(context.TODO(), concert.GetAvailableSpec{Offset: 0, Limit: 10, Search: "Test"})
		assert.NoError(t, err)
		assert.Equal(t, total, expTotal)
		assert.NotEmpty(t, results)
	})
}
