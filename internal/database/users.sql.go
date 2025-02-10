// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, hashed_password, first_name, last_name)
VALUES (
	?,
	?,
	?,
	?
)
RETURNING id, email, hashed_password, first_name, last_name, created_at, updated_at, is_admin
`

type CreateUserParams struct {
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Email,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const getUserAdmin = `-- name: GetUserAdmin :one
SELECT 
	id,
	email,
	first_name,
	last_name,
	created_at,
	updated_at,
	is_admin
FROM users
WHERE id = ?
`

type GetUserAdminRow struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	IsAdmin   sql.NullInt64
}

func (q *Queries) GetUserAdmin(ctx context.Context, id int64) (GetUserAdminRow, error) {
	row := q.db.QueryRowContext(ctx, getUserAdmin, id)
	var i GetUserAdminRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT 
	email, 
	first_name, 
	last_name, 
	created_at, 
	updated_at
FROM users
WHERE id = ?
`

type GetUserByIDRow struct {
	Email     string
	FirstName string
	LastName  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) GetUserByID(ctx context.Context, id int64) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsersAdmin = `-- name: ListUsersAdmin :many

SELECT 
	id,
	email,
	first_name,
	last_name,
	created_at,
	updated_at,
	is_admin
FROM users
`

type ListUsersAdminRow struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	IsAdmin   sql.NullInt64
}

// -------------------------------
// ADMIN
// -------------------------------
func (q *Queries) ListUsersAdmin(ctx context.Context) ([]ListUsersAdminRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsersAdmin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersAdminRow
	for rows.Next() {
		var i ListUsersAdminRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsAdmin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users 
SET 
	email = ?,
	hashed_password = ?,
	first_name = ?,
	last_name = ?,
	updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserParams struct {
	Email          string
	HashedPassword string
	FirstName      string
	LastName       string
	ID             int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Email,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.ID,
	)
	return err
}

const updateUserAdmin = `-- name: UpdateUserAdmin :exec
UPDATE users
SET 
	email = ?,
	first_name = ?,
	last_name = ?,
	updated_at = CURRENT_TIMESTAMP,
	is_admin = ?
WHERE id = ?
`

type UpdateUserAdminParams struct {
	Email     string
	FirstName string
	LastName  string
	IsAdmin   sql.NullInt64
	ID        int64
}

func (q *Queries) UpdateUserAdmin(ctx context.Context, arg UpdateUserAdminParams) error {
	_, err := q.db.ExecContext(ctx, updateUserAdmin,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.IsAdmin,
		arg.ID,
	)
	return err
}
