-- name: ListUserRentals :many
SELECT 
		start_time,
		end_time,
		start_latitude,
		start_longitude,
		end_latitude,
		end_longitude
FROM rentals
WHERE user_id = ?;

-- name: CreateRental :one
INSERT INTO rentals (user_id, bike_id, status, start_latitude, start_longitude)
VALUES (
	?,
	?,
	'running',
	?,
	?
)
RETURNING *;

-- name: UpdateRentalToEnded :exec
UPDATE rentals 
SET 
	status = 'ended',
	end_latitude = ?,
	end_longitude = ?
WHERE user_id = ? AND status = 'running'; -- the constraint in unique index for one running rental per user allows this query

---------------------------------
-- ADMIN
---------------------------------

-- name: ListRentalsAdmin :many
SELECT 
	rentals.id AS rental_id,
	rentals.bike_id,
	users.email AS user_email,
	users.first_name AS user_first_name,
	users.last_name AS user_last_name,
	rentals.status,
	rentals.start_time,
	rentals.end_time,
	rentals.start_latitude,
	rentals.start_longitude,
	rentals.end_latitude,
	rentals.end_longitude
FROM rentals
INNER JOIN users
	ON rentals.user_id = users.id
ORDER BY rentals.start_time DESC; 

-- name: GetRentalAdmin :one
SELECT 
	rentals.bike_id,
	users.email AS user_email,
	users.first_name AS user_first_name,
	users.last_name AS user_last_name,
	rentals.status,
	rentals.start_time,
	rentals.end_time,
	rentals.start_latitude,
	rentals.start_longitude,
	rentals.end_latitude,
	rentals.end_longitude
FROM rentals
INNER JOIN users
	ON rentals.user_id = users.id
WHERE rentals.id = ?;

-- name: UpdateRentalAdmin :exec
UPDATE rentals
SET 
	user_id = ?,
	bike_id = ?,
	status = ?,
	start_time = ?,
	end_time = ?,
	start_latitude = ?,
	start_longitude = ?,
	end_latitude = ?,
	end_longitude = ?
WHERE id = ?;

