-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key)
    VALUES ($1, $2, $3, $4, encode(sha256 (random()::text::bytea), 'hex'))
RETURNING
    *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    name = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;
