-- name: CreateMail :one
INSERT INTO mails (
 title, body, receiver, sender 
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ListMails :many
SELECT * FROM mails
WHERE receiver = $1;

-- name: GetMail :one
SELECT * FROM mails
WHERE id = $1 LIMIT 1;

-- name: UpdateMail :one
UPDATE mails SET title = $2, body = $3
WHERE id = $1
RETURNING *;

-- name: DeleteMail :exec
DELETE FROM mails WHERE id = $1;