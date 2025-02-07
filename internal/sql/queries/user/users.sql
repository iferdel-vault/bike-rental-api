-- name: CreateUser :one
INSERT INTO users (email, hashed_password, first_name, last_name)
VALUES (
	?,
	?,
	?,
	?
)
RETURNING *;

-- name: GetUserByID :one
SELECT 
	email, 
	first_name, 
	last_name, 
	created_at, 
	updated_at
FROM users
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users 
SET 
	email = ?,
	hashed_password = ?,
	first_name = ?,
	last_name = ?,
	updated_at = CURRENT_TIMESTAMP
WHERE id = ?;
