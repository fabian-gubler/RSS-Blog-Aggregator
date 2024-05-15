-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name text NOT NULL
);

-- +goose Down
DROP TABLE users;
