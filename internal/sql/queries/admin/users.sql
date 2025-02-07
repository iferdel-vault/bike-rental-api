-- name: ListUsers :many
SELECT 
	id,
	email,
	first_name,
	last_name,
	created_at,
	updated_at,
	is_admin
FROM users;

-- name: GetUser :one
SELECT 
	id,
	email,
	first_name,
	last_name,
	created_at,
	updated_at,
	is_admin
FROM users
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users
SET 
	email = ?,
	first_name = ?,
	last_name = ?,
	updated_at = CURRENT_TIMESTAMP,
	is_admin = ?
WHERE id = ?;

