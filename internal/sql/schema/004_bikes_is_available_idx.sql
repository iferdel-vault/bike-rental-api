-- +goose Up
CREATE INDEX bikes_is_available_idx ON bikes (is_available);

-- +goose Down
DROP INDEX bikes_is_available_idx;
