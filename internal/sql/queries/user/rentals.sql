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
