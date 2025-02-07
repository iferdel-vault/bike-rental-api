-- +goose Up
CREATE TABLE rentals (
    id INTEGER PRIMARY KEY, -- sqlite does not have uuid by default
		user_id INTEGER NOT NULL,
		bike_id INTEGER NOT NULL,
		status TEXT NOT NULL,
		start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- sqlite does not have more than text, integer and real. text is more readable and may fit better here.
		end_time TIMESTAMP,
		start_latitude REAL CHECK (start_latitude BETWEEN -90 AND 90),
		start_longitude REAL CHECK (start_longitude BETWEEN -180 AND 180),
		end_latitude REAL CHECK (end_latitude BETWEEN -90 AND 90),
		end_longitude REAL CHECK (end_longitude BETWEEN -180 AND 180),
		CONSTRAINT fk_users
			FOREIGN KEY (user_id)
			REFERENCES users(id),
		CONSTRAINT fk_bikes
			FOREIGN KEY (bike_id)
			REFERENCES bikes(id)
);

-- +goose Down
DROP TABLE rentals;
