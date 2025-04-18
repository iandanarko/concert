package concert

import "time"

type Concert struct {
	ID        uint64
	Name      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
