-- name: CreateCurrency :one
INSERT INTO currencies (
    name,
    symbol,
    isdefault,
    isdeleted,
    isblocked
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetCurrency :one
SELECT * FROM currencies WHERE id = $1;

-- name: GetAllCurrency :many
SELECT * FROM currencies;

-- name: UpdateCurrency :one
UPDATE currencies
SET
    name = $2,
    symbol = $3,
    isdefault = $4,
    isdeleted = $5,
    isblocked = $6
WHERE id = $1
RETURNING *;

-- name: DeleteCurrency :one
DELETE FROM currencies WHERE id = $1
RETURNING *;
