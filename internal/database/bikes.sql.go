// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: bikes.sql

package database

import (
	"context"
	"database/sql"
)

const createBikeAdmin = `-- name: CreateBikeAdmin :one
INSERT INTO bikes (is_available, latitude, longitude)
VALUES (
	1,
	?,
	?
)
RETURNING id, is_available, latitude, longitude, created_at, updated_at
`

type CreateBikeAdminParams struct {
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
}

func (q *Queries) CreateBikeAdmin(ctx context.Context, arg CreateBikeAdminParams) (Bike, error) {
	row := q.db.QueryRowContext(ctx, createBikeAdmin, arg.Latitude, arg.Longitude)
	var i Bike
	err := row.Scan(
		&i.ID,
		&i.IsAvailable,
		&i.Latitude,
		&i.Longitude,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBikesAdmin = `-- name: ListBikesAdmin :many

SELECT 
	id,
	is_available,
	latitude,
	longitude,
	created_at,
	updated_at
FROM bikes
`

// -------------------------------
// ADMIN
// -------------------------------
func (q *Queries) ListBikesAdmin(ctx context.Context) ([]Bike, error) {
	rows, err := q.db.QueryContext(ctx, listBikesAdmin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bike
	for rows.Next() {
		var i Bike
		if err := rows.Scan(
			&i.ID,
			&i.IsAvailable,
			&i.Latitude,
			&i.Longitude,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listBikesAvailable = `-- name: ListBikesAvailable :many

SELECT 
	latitude, 
	longitude 
FROM bikes
WHERE is_available
`

type ListBikesAvailableRow struct {
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
}

// -------------------------------
// USER
// -------------------------------
func (q *Queries) ListBikesAvailable(ctx context.Context) ([]ListBikesAvailableRow, error) {
	rows, err := q.db.QueryContext(ctx, listBikesAvailable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListBikesAvailableRow
	for rows.Next() {
		var i ListBikesAvailableRow
		if err := rows.Scan(&i.Latitude, &i.Longitude); err != nil {
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

const updateBikeToAvailableAdmin = `-- name: UpdateBikeToAvailableAdmin :exec
UPDATE bikes 
SET 
	is_available = 1,
	latitude = ?, -- way to update tracking of geolocation
	longitude = ?, -- end values for geolocation after ending a rental
	updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND is_available = 0
`

type UpdateBikeToAvailableAdminParams struct {
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
	ID        int64
}

func (q *Queries) UpdateBikeToAvailableAdmin(ctx context.Context, arg UpdateBikeToAvailableAdminParams) error {
	_, err := q.db.ExecContext(ctx, updateBikeToAvailableAdmin, arg.Latitude, arg.Longitude, arg.ID)
	return err
}

const updateBikeToUnavailableAdmin = `-- name: UpdateBikeToUnavailableAdmin :exec
UPDATE bikes 
SET 
	is_available = 0,
	latitude = ?, -- start values of geolocation for a create rental
	longitude = ?,
	updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND is_available = 1
`

type UpdateBikeToUnavailableAdminParams struct {
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
	ID        int64
}

func (q *Queries) UpdateBikeToUnavailableAdmin(ctx context.Context, arg UpdateBikeToUnavailableAdminParams) error {
	_, err := q.db.ExecContext(ctx, updateBikeToUnavailableAdmin, arg.Latitude, arg.Longitude, arg.ID)
	return err
}
