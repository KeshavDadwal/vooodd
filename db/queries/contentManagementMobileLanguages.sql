-- name: CreateContentManagementMobileLanguage :one
INSERT INTO content_management_mobile_languages (content_management_mobile_id, language_id, title, link, link_text, content)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: SelectContentManagementMobileLanguages :one
SELECT * FROM content_management_mobile_languages WHERE id = $1;

-- name: SelectAllContentManagementMobileLanguages :many
SELECT * FROM content_management_mobile_languages;

-- name: UpdateContentManagementMobileLanguage :one
UPDATE content_management_mobile_languages
SET title = $3, link = $4, link_text = $5, content = $6
WHERE id = $1 AND content_management_mobile_id = $2
RETURNING id;

-- name: DeleteContentManagementMobileLanguage :one
DELETE FROM content_management_mobile_languages WHERE id = $1
RETURNING id;
