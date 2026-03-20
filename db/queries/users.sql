-- name: GetUserByID :one
SELECT id, email, username FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, username FROM users ORDER BY id
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (email, pass, username) VALUES ($1, $2, $3)
RETURNING id, email, username;

-- name: EditUser :one
UPDATE users SET
email = COALESCE($2, email),
username = COALESCE($3, username)
WHERE id = $1
RETURNING id, email, username;

-- name: EditUserPass :one
UPDATE users SET
pass = $2
WHERE id = $1
RETURNING id, email, username;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
