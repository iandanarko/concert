-- +goose Up
-- +goose StatementBegin
CREATE TABLE bookings (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  concert_id BIGINT NOT NULL,
  ticket_window_id BIGINT NOT NULL,
  ticket_count INT NOT NULL,
  created_at DATETIME NOT NULL,
  CONSTRAINT fk_concert_booking FOREIGN KEY (concert_id) REFERENCES concerts(id),
  CONSTRAINT fk_window_booking FOREIGN KEY (ticket_window_id) REFERENCES ticket_windows(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bookings;
-- +goose StatementEnd
