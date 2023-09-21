-- name: CreateLanguage :one
INSERT INTO languages (
    name,
    code,
    isdefault,
    isdeleted,
    isblocked
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
        )
RETURNING *;

-- name: GetAllLanguage :many
SELECT * FROM languages;

-- name: UpdateLanguage :one
UPDATE languages
SET name = $1, code = $2, isdefault = $3, isdeleted = $4, isblocked = $5
WHERE id = $6
RETURNING *;

-- name: DeleteLanguage :one
DELETE FROM languages WHERE id = $1
RETURNING *;
