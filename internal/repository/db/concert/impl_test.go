package concertrepo_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	concertrepo "github.com/iandanarko/concert/internal/repository/db/concert"
)

type implTest struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
	repo concertrepo.Impl
}

func new(t *testing.T) implTest {
	db, mock, _ := sqlmock.New()
	repo := concertrepo.New(db)
	return implTest{
		db:   db,
		mock: mock,
		repo: repo,
	}
}
