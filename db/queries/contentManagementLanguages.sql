-- name: CreateContentManagementLanguage :one
INSERT INTO content_management_languages (content_management_id, language_id, name, title, content)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: SelectContentManagementLanguages :one
SELECT * FROM content_management_languages WHERE id = $1;


-- name: SelectAllContentManagementLanguages :many
SELECT * FROM content_management_languages;

-- name: UpdateContentManagementLanguage :one
UPDATE content_management_languages
SET name = $3, title = $4, content = $5
WHERE id = $1 AND content_management_id = $2
RETURNING id;

-- name: DeleteContentManagementLanguage :one
DELETE FROM content_management_languages WHERE id = $1
RETURNING id;
