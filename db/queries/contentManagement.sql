-- name: CreateContentManagement :one
INSERT INTO content_managements (name, title, content, isdeleted, isblocked)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: SelectContentManagements :one
SELECT * FROM content_managements WHERE id = $1;

-- name: SelectAllContentManagements :many
SELECT * FROM content_managements;

-- name: UpdateContentManagement :one
UPDATE content_managements
SET name = $2, title = $3, content = $4, isdeleted = $5, isblocked = $6
WHERE id = $1
RETURNING id;

-- name: DeleteContentManagement :one
DELETE FROM content_managements WHERE id = $1
RETURNING id;
