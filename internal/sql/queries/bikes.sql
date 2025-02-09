---------------------------------
-- USER
---------------------------------

-- name: ListBikesAvailable :many
SELECT 
	latitude, 
	longitude 
FROM bikes
WHERE is_available;


---------------------------------
-- ADMIN
---------------------------------

-- name: ListBikesAdmin :many
SELECT 
	id,
	is_available,
	latitude,
	longitude,
	created_at,
	updated_at
FROM bikes;

-- name: CreateBikeAdmin :one
INSERT INTO bikes (is_available, latitude, longitude)
VALUES (
	1,
	?,
	?
)
RETURNING *;

-- name: UpdateBikeToAvailableAdmin :exec
UPDATE bikes 
SET 
	is_available = 1,
	latitude = ?, -- way to update tracking of geolocation
	longitude = ?, -- end values for geolocation after ending a rental
	updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND is_available = 0;

-- name: UpdateBikeToUnavailableAdmin :exec
UPDATE bikes 
SET 
	is_available = 0,
	latitude = ?, -- start values of geolocation for a create rental
	longitude = ?,
	updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND is_available = 1;

