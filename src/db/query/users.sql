-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByToken :one
SELECT * FROM users
WHERE token = $1
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (email, password, token, updated_at)
VALUES($1, $2, $3, now());

-- name: UpdateUserToken :exec
UPDATE users SET token = $2, updated_at = now()
WHERE token = $1;

-- name: UpdateUserPassword :exec
UPDATE users SET password = $2, updated_at = now()
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET email = $2, password = $3, token = $4, updated_at = now()
WHERE id = $1;
