-- internal/auth/infrastructure/repository/query.sql

-- name: CreateUser :one
INSERT INTO users (id, email, password_hash, role, verified, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: UpdateUserVerified :exec
UPDATE users SET verified = $2, updated_at = $3 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
