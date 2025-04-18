package bookingrepo

import (
	"context"
	"database/sql"

	cErr "github.com/iandanarko/concert/internal/error"
	"github.com/iandanarko/concert/internal/model/booking"
)

func (i Impl) Book(ctx context.Context, spec booking.BookSpec) error {
	tx, err := i.db.Begin()
	if err != nil {
		return err
	}

	query := `SELECT id, ticket_limit, tickets_sold
		FROM ticket_windows
		WHERE concert_id = ? AND NOW() BETWEEN window_start AND window_end FOR UPDATE`

	var ticketWindow booking.TicketWindow
	err = tx.QueryRowContext(ctx, query, spec.ConcertID).Scan(&ticketWindow.ID, &ticketWindow.TicketLimit, &ticketWindow.TicketSold)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return err
	}

	if err == sql.ErrNoRows {
		tx.Rollback()
		return cErr.ErrTicketWindowNotFound
	}

	if ticketWindow.TicketLimit < ticketWindow.TicketSold+spec.NumTickets {
		tx.Rollback()
		return cErr.ErrOutOfTickets
	}

	queryBooking := `INSERT INTO bookings (user_id, concert_id, ticket_window_id, ticket_count, created_at) 
		VALUES (?, ?, ?, ?, NOW())`

	_, err = tx.ExecContext(ctx, queryBooking, spec.UserID, spec.ConcertID, ticketWindow.ID, spec.NumTickets)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryUpdate := `UPDATE ticket_windows SET tickets_sold = tickets_sold + ?, updated_at = NOW() WHERE id = ?`
	_, err = tx.ExecContext(ctx, queryUpdate, spec.NumTickets, ticketWindow.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
