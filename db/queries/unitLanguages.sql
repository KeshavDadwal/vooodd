-- name: CreateUnitLanguage :one
INSERT INTO unit_languages (
    unit_id,
    language_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectUnitLanguages :many
SELECT * FROM unit_languages;

-- name: UpdateUnitLanguage :one
UPDATE unit_languages
SET
    unit_id = $2,
    language_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteUnitLanguage :one
DELETE FROM unit_languages WHERE id = $1
RETURNING id;
