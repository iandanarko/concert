-- +goose Up
-- +goose StatementBegin
CREATE INDEX date_name_on_concerts ON concerts (date, (LOWER(name)));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX date_name_on_concerts ON concerts;
-- +goose StatementEnd
