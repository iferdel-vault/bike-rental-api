-- +goose Up
CREATE UNIQUE INDEX rentals_bike_active_idx ON rentals (bike_id)
	WHERE status = 'running';

-- +goose Down
DROP INDEX rentals_bike_active_idx;
