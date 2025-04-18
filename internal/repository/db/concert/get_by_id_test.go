package concertrepo_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConcert_GetByID(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT id, name, date, created_at, updated_at 
		FROM concerts
		WHERE id = ?`)
	t.Run("Failed: error query", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()

		suite.mock.ExpectQuery(query).WillReturnError(errors.New("some error"))
		result, err := suite.repo.GetByID(context.TODO(), 1)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("Failed: error not found", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()

		suite.mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)
		result, err := suite.repo.GetByID(context.TODO(), 1)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("Success: success get", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()

		suite.mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "date", "created_at", "updated_at"}).
			AddRow(1, "Concert test", time.Now(), time.Now(), time.Now()))
		result, err := suite.repo.GetByID(context.TODO(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}
