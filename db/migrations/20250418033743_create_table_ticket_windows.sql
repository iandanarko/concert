-- +goose Up
-- +goose StatementBegin
CREATE TABLE ticket_windows (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  concert_id BIGINT NOT NULL,
  window_start DATETIME NOT NULL,
  window_end DATETIME NOT NULL,
  ticket_limit INT NOT NULL,
  tickets_sold INT DEFAULT 0,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  CONSTRAINT fk_concert FOREIGN KEY (concert_id) REFERENCES concerts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ticket_windows;
-- +goose StatementEnd
