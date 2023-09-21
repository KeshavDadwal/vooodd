-- name: CreateHomePageIconLanguage :one
INSERT INTO home_page_icon_languages (home_page_icon_id, language_id, name)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectAllHomePageIconLanguages :many
SELECT * FROM home_page_icon_languages;

-- name: UpdateHomePageIconLanguage :one
UPDATE home_page_icon_languages
SET
    home_page_icon_id = $2,
    language_id = $3,
    name = $4
WHERE id = $1
RETURNING id;

-- name: DeleteHomePageIconLanguage :one
DELETE FROM home_page_icon_languages WHERE id = $1
RETURNING id;
