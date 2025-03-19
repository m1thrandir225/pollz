-- name: CreateVote :one
INSERT INTO votes(
    option_id,
    ip_address,
    user_agent
)
VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetVote :one
SELECT *
FROM votes
WHERE id=$1
LIMIT 1;

-- name: UpdateVoteOption :one
UPDATE votes
SET option_id=$2
WHERE id=$1
RETURNING *;

-- name: DeleteVote :one
DELETE FROM votes
WHERE id=$1
RETURNING *;
