-- name: CreatePollOption :one
INSERT INTO poll_options(
    poll_id,
    option_text
) VALUES (
    sqlc.arg(poll_id)::uuid,
    sqlc.arg(option_text)::text
) RETURNING *;

-- name: DeletePollOption :one
DELETE from poll_options
WHERE id = sqlc.arg(id)::uuid
RETURNING *;