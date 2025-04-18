package bookingrepo_test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	bookingrepo "github.com/iandanarko/concert/internal/repository/db/booking"
)

type implTest struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo bookingrepo.Impl
}

func new() implTest {
	db, mock, _ := sqlmock.New()
	repo := bookingrepo.New(db)
	return implTest{
		db:   db,
		mock: mock,
		repo: repo,
	}
}
