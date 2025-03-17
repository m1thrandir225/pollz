-- name: CreatePoll :one
INSERT INTO polls(
    description,
    created_by,
    active_until
)
VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetPolls :many
SELECT * FROM polls;

-- name: GetPoll :one
SELECT * FROM polls
WHERE id = $1
LIMIT 1;

-- name: GetPollWithOptions :one
SELECT * FROM polls as p
JOIN poll_options as po ON p.id = po.poll_id
WHERE p.id = $1;

-- name: UpdatePoll :one
UPDATE polls
SET description=$2, active_until=$3
WHERE id = $1
RETURNING  *;

-- name: UpdatePollStatus :one
UPDATE polls
SET active_until=$2
WHERE id = $1
RETURNING *;

-- name: DeletePoll :one
DELETE FROM polls
WHERE id = $1
RETURNING  *;

