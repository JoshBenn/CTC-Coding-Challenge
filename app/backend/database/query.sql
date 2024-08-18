-- Users
-- name: Createuser :one
INSERT INTO
    users (
        email,
        username,
        password
    )
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;
