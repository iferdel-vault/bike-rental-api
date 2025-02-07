-- name: ListBikesAvailable :many
SELECT 
	latitude, 
	longitude 
FROM bikes
WHERE is_available;

