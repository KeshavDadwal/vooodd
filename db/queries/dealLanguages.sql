-- name: CreateDealsLanguage :one
INSERT INTO deals_languages (deals_id, language_id, name)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectDealsLanguages :one
SELECT * FROM deals_languages where id = $1;

-- name: SelectAllDealsLanguages :many
SELECT * FROM deals_languages;


-- name: UpdateDealsLanguage :one
UPDATE deals_languages
SET
    deals_id = $2,
    language_id = $3,
    name = $4
WHERE id = $1
RETURNING id;

-- name: DeleteDealsLanguage :one
DELETE FROM deals_languages WHERE id = $1
RETURNING id;
