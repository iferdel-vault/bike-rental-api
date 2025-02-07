-- +goose Up
ALTER TABLE rentals
ADD COLUMN duration INTEGER GENERATED ALWAYS AS (
	ceil((strftime('%s', end_time) - strftime('%s', start_time)) / 60)
) VIRTUAL;

-- +goose Down
ALTER TABLE rentals
DROP COLUMN duration;
