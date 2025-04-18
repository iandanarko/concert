package bookingrepo

import "database/sql"

type Impl struct {
	db *sql.DB
}

func New(db *sql.DB) Impl {
	return Impl{
		db: db,
	}
}
