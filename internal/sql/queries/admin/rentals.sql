-- name: ListRentals :many
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

-- name: GetRental :one
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

-- name: UpdateRental :exec
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

