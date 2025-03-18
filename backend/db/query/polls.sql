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
SELECT * FROM polls
WHERE created_by = $1;

-- name: GetPoll :one
SELECT * FROM polls
WHERE id = $1
LIMIT 1;

-- name: IsPollActive :one
SELECT * FROM polls as p 
WHERE p.active_until > now() AND p.id = $1
LIMIT 1;

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

