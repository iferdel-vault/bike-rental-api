-- +goose Up
CREATE UNIQUE INDEX rentals_user_active_idx ON rentals (user_id)
	WHERE status = 'running';

-- +goose Down
DROP INDEX rentals_user_active_idx;
