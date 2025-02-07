-- +goose Up
CREATE TABLE bikes (
    id INTEGER PRIMARY KEY,
		is_available INTEGER, -- if its not available it may be due to being rented at the time or being, for example, in maintenance
		latitude REAL,
		longitude REAL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE bikes;
