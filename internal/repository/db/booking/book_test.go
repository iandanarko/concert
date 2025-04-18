package bookingrepo_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/iandanarko/concert/internal/model/booking"
	"github.com/stretchr/testify/assert"
)

func TestBooking_Book(t *testing.T) {
	queryGet := regexp.QuoteMeta(`SELECT id, ticket_limit, tickets_sold
		FROM ticket_windows
		WHERE concert_id = ? AND NOW() BETWEEN window_start AND window_end FOR UPDATE`)
	queryBook := regexp.QuoteMeta(`INSERT INTO bookings (user_id, concert_id, ticket_window_id, ticket_count, created_at) 
		VALUES (?, ?, ?, ?, NOW())`)
	queryUpdate := regexp.QuoteMeta(`UPDATE ticket_windows SET tickets_sold = tickets_sold + ?, updated_at = NOW() WHERE id = ?`)

	t.Run("Failed: error begin", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin().WillReturnError(errors.New("some error"))
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Error(t, err)
	})

	t.Run("Failed: error find ticket window", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(queryGet).WillReturnError(errors.New("some error"))
		suite.mock.ExpectRollback()
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Error(t, err)
	})

	t.Run("Failed: ticket window not found", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(queryGet).WillReturnError(sql.ErrNoRows)
		suite.mock.ExpectRollback()
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Error(t, err)
	})

	t.Run("Failed: out of stocks", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(queryGet).WillReturnRows(sqlmock.NewRows([]string{"id", "ticket_limit", "ticker_sold"}).
			AddRow(1, 10, 10))
		suite.mock.ExpectRollback()
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Error(t, err)
	})

	t.Run("Failed: failed insert", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(queryGet).WillReturnRows(sqlmock.NewRows([]string{"id", "ticket_limit", "ticker_sold"}).
			AddRow(1, 10, 0))
		suite.mock.ExpectExec(queryBook).WillReturnError(errors.New("some error"))
		suite.mock.ExpectRollback()
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Error(t, err)
	})

	t.Run("Failed: failed update", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(queryGet).WillReturnRows(sqlmock.NewRows([]string{"id", "ticket_limit", "ticker_sold"}).
			AddRow(1, 10, 0))
		suite.mock.ExpectExec(queryBook).WillReturnResult(sqlmock.NewResult(1, 1))
		suite.mock.ExpectExec(queryUpdate).WillReturnError(errors.New("some error"))
		suite.mock.ExpectRollback()
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Error(t, err)
	})

	t.Run("Success: success book", func(t *testing.T) {
		t.Parallel()
		suite := new()
		defer suite.db.Close()
		suite.mock.ExpectBegin()
		suite.mock.ExpectQuery(queryGet).WillReturnRows(sqlmock.NewRows([]string{"id", "ticket_limit", "ticker_sold"}).
			AddRow(1, 10, 0))
		suite.mock.ExpectExec(queryBook).WillReturnResult(sqlmock.NewResult(1, 1))
		suite.mock.ExpectExec(queryUpdate).WillReturnResult(sqlmock.NewResult(1, 1))
		suite.mock.ExpectCommit()
		err := suite.repo.Book(context.TODO(), booking.BookSpec{NumTickets: 1})
		assert.Nil(t, err)
	})
}
