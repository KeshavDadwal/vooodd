-- name: CreateContentManagementBannerLanguage :one
INSERT INTO content_management_banner_languages (content_management_id, language_id, name, title, content)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: SelectContentManagementBannerLanguages :one
SELECT * FROM content_management_banner_languages WHERE  id = $1;

-- name: SelectAllContentManagementBannerLanguages :many
SELECT * FROM content_management_banner_languages;


-- name: UpdateContentManagementBannerLanguage :one
UPDATE content_management_banner_languages
SET name = $3, title = $4, content = $5
WHERE id = $1 AND content_management_id = $2
RETURNING id;

-- name: DeleteContentManagementBannerLanguage :one
DELETE FROM content_management_banner_languages WHERE id = $1
RETURNING id;
