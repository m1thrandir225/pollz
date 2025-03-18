-- name: CreatePollOption :one
INSERT INTO poll_options(
    poll_id,
    option_text
) VALUES (
    sqlc.arg(poll_id)::uuid,
    sqlc.arg(option_text)::text
) RETURNING *;


-- name: CreateMultipleOptions :many 
INSERT INTO poll_options(
  poll_id,
  option_text
) VALUES (
  sqlc.arg(poll_id)::uuid,
  unnest(sqlc.arg(option_texts)::text[])
) RETURNING *;

-- name: GetOption :one
SELECT * FROM poll_options
WHERE id=$1
LIMIT 1;

-- name: GetOptionsForPoll :many 
SELECT * FROM poll_options
WHERE poll_id = $1;

-- name: UpdatePollOption :one
UPDATE poll_options
SET option_text=$2
WHERE id=$1
RETURNING *;

-- name: DeletePollOption :one
DELETE from poll_options
WHERE id = sqlc.arg(id)::uuid
RETURNING *;
