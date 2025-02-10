-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		hashed_password TEXT NOT NULL DEFAULT 'unset',
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP
	);

-- +goose Down
DROP TABLE users;
