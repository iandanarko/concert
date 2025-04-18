package concertrepo

import (
	"context"
	"database/sql"

	"github.com/iandanarko/concert/internal/model/concert"
)

func (i Impl) GetByID(ctx context.Context, id uint64) (*concert.Concert, error) {
	query := `
		SELECT id, name, date, created_at, updated_at 
		FROM concerts
		WHERE id = ?`

	var result concert.Concert
	err := i.db.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.Name,
		&result.Date,
		&result.CreatedAt,
		&result.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, nil
}
