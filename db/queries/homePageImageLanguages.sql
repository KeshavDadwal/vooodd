-- name: CreateHomePageImageLanguage :one
INSERT INTO home_page_image_languages (home_page_image_id, language_id, name)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectHomePageImageLanguages :many
SELECT * FROM home_page_image_languages;

-- name: UpdateHomePageImageLanguage :one
UPDATE home_page_image_languages
SET
    home_page_image_id = $2,
    language_id = $3,
    name = $4
WHERE id = $1
RETURNING id;

-- name: DeleteHomePageImageLanguage :one
DELETE FROM home_page_image_languages WHERE id = $1
RETURNING id;
