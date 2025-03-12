-- name: CreateUser :one
INSERT INTO users(
    first_name,
    last_name,
    email,
    password
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING id, first_name, last_name, email, created_at;

-- name: GetUserDetails :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id, first_name, last_name, email, created_at;