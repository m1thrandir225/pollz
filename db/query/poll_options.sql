-- name: CreatePollOption :one
INSERT INTO poll_options(
    poll_id,
    option_text
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: DeletePollOption :one
DELETE from poll_options
WHERE id = $1
RETURNING *;