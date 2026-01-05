-- +goose Up
ALTER TABLE users
ADD COLUMN is_chirpy_red BOOLEAN NOT NULL
DEFAULT FALSE;

-- +goose Down
ALTER TABLE user
DROP COLUMN is_chirpy_red;