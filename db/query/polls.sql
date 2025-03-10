-- name: CreatePoll :one
INSERT INTO polls(
    description,
    created_by
)
VALUES (
    $1,
    $2
) RETURNING id, description, created_by, is_active, created_at;

-- name: DisablePoll :one
UPDATE polls
SET is_active=$2
WHERE id = $1
RETURNING *;

