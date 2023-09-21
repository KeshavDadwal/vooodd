-- name: CreateMallLanguage :one
INSERT INTO mall_languages (
    mall_id,
    language_id,
    name,
    address,
    city,
    state,
    zip,
    phone_no
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING id;

-- name: SelectMallLanguages :many
SELECT * FROM mall_languages;

-- name: UpdateMallLanguage :one
UPDATE mall_languages
SET
    mall_id = $2,
    language_id = $3,
    name = $4,
    address = $5,
    city = $6,
    state = $7,
    zip = $8,
    phone_no = $9
WHERE id = $1
RETURNING *;


-- name: DeleteMallLanguage :one
DELETE FROM mall_languages WHERE id = $1
RETURNING id;

