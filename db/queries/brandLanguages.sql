-- name: CreateBrandLanguage :one
INSERT INTO brands_languages (brand_id, language_id, name)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectBrandLanguages :one
SELECT * FROM brands_languages WHERE id = $1;

-- name: SelectALLBrandLanguages :many
SELECT * FROM brands_languages;

-- name: UpdateBrandLanguage :one
UPDATE brands_languages
SET name = $3
WHERE id = $1 AND brand_id = $2
RETURNING id;

-- name: DeleteBrandLanguage :one
DELETE FROM brands_languages WHERE id = $1
RETURNING id;
