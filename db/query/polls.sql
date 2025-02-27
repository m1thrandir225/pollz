-- name: CreatePoll :one
INSERT INTO polls(
    question,
    created_by
)
VALUES (
    $1,
    $2
) RETURNING id, question, created_by, is_active, created_at;

-- name: DisablePoll :one
UPDATE polls
SET is_active=$2
WHERE id = $1
RETURNING *;

