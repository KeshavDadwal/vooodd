-- name: CreateStoreLanguage :one
INSERT INTO store_languages (
    language_id,
    store_id,
    name,
    description
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: SelectStoreLanguages :many
SELECT * FROM store_languages;

-- name: UpdateStoreLanguage :one
UPDATE store_languages
SET
    language_id = $2,
    store_id = $3,
    name = $4,
    description = $5
WHERE id = $1
RETURNING id;

-- name: DeleteStoreLanguage :one
DELETE FROM store_languages WHERE id = $1
RETURNING id;
