package concertrepo

import (
	"context"
	"database/sql"

	"github.com/iandanarko/concert/internal/model/concert"
)

func (i Impl) GetAvailableConcerts(ctx context.Context, spec concert.GetAvailableSpec) ([]concert.Concert, error) {
	args := []any{}
	cond := "date > now()"
	if spec.Search != "" {
		cond += " AND name LIKE (?%)"
		args = append(args, spec.Search)
	}

	query := `
		SELECT id, name, date, created_at, updated_at 
		FROM concerts 
		WHERE ` + cond + `
		ORDER BY date ASC
	`

	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return []concert.Concert{}, err
	}

	result := []concert.Concert{}
	for rows.Next() {
		c := concert.Concert{}
		err := rows.Scan(&c.ID, &c.Name, &c.Date, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return []concert.Concert{}, err
		}
		result = append(result, c)
	}

	return result, nil
}
