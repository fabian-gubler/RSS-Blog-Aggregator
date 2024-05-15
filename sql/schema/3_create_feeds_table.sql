-- +goose Up
CREATE TABLE IF NOT EXISTS feeds (
    id uuid PRIMARY KEY,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name text NOT NULL,
    url text NOT NULL UNIQUE,
    user_id uuid NOT NULL
    -- TODO: Add a foreign key constraint to the user_id column
    -- CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
