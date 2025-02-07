-- +goose Up
CREATE INDEX rentals_status_idx ON rentals (status);

-- +goose Down
DROP INDEX rentals_status_idx;
