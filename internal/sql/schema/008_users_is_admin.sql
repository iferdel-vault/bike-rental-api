-- +goose Up
ALTER TABLE users
ADD COLUMN is_admin INTEGER DEFAULT 0;

-- +goose Down
ALTER TABLE users
DROP COLUMN is_admin;
