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

---------------------------------
-- ADMIN
---------------------------------

-- name: ListUsersAdmin :many
SELECT 
	id,
	email,
	first_name,
	last_name,
	created_at,
	updated_at,
	is_admin
FROM users;

-- name: GetUserAdmin :one
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

-- name: UpdateUserAdmin :exec
UPDATE users
SET 
	email = ?,
	first_name = ?,
	last_name = ?,
	updated_at = CURRENT_TIMESTAMP,
	is_admin = ?
WHERE id = ?;

