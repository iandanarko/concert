-- +goose Up
-- +goose StatementBegin
CREATE INDEX concert_id_window_start_window_end_on_ticket_windows ON ticket_windows (concert_id, window_start, window_end);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX concert_id_window_start_window_end_on_ticket_windows ON ticket_windows;
-- +goose StatementEnd
