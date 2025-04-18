package concertrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iandanarko/concert/internal/model/concert"
)

func (i Impl) GetAvailableConcerts(ctx context.Context, spec concert.GetAvailableSpec) ([]concert.Concert, uint64, error) {
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
		OFFSET ` + fmt.Sprint(spec.Offset) + ` LIMIT ` + fmt.Sprint(spec.GetLimit()) + `
	`

	cntQuery := `SELECT COUNT(*) FROM concerts WHERE ` + cond

	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return []concert.Concert{}, 0, err
	}

	result := []concert.Concert{}
	for rows.Next() {
		c := concert.Concert{}
		err := rows.Scan(&c.ID, &c.Name, &c.Date, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return []concert.Concert{}, 0, err
		}
		result = append(result, c)
	}

	var total uint64
	err = i.db.QueryRowContext(ctx, cntQuery, args...).Scan(&total)
	if err != nil {
		return []concert.Concert{}, 0, err
	}

	return result, total, nil
}
