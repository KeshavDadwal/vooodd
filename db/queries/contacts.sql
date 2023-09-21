-- name: CreateContact :one
INSERT INTO contacts (name, email, user_id, subject, message)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: SelectContacts :one
SELECT * FROM contacts where id =$1;

-- name: SelectAllContacts :many
SELECT * FROM contacts;

-- name: UpdateContact :one
UPDATE contacts
SET name = $2, email = $3, user_id = $4, subject = $5, message = $6
WHERE id = $1
RETURNING id;

-- name: DeleteContact :one
DELETE FROM contacts WHERE id = $1
RETURNING id;
