-- name: CreateCompetitorLanguage :one
INSERT INTO competitor_languages (competitor_id, language_id, name)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectCompetitorLanguages :one
SELECT * FROM competitor_languages WHERE id = $1;

-- name: SelectAllCompetitorLanguages :many
SELECT * FROM competitor_languages;

-- name: UpdateCompetitorLanguage :one
UPDATE competitor_languages
SET name = $3
WHERE id = $1 AND competitor_id = $2
RETURNING id;

-- name: DeleteCompetitorLanguage :one
DELETE FROM competitor_languages WHERE id = $1
RETURNING id;
