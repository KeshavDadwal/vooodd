-- name: CreateUnit :one
INSERT INTO units (name, isdeleted, isblocked)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectUnits :many
SELECT * FROM units WHERE isdeleted = false AND isblocked = false;

-- name: UpdateUnit :one
UPDATE units
SET name = $2, isdeleted = $3, isblocked = $4
WHERE id = $1
RETURNING id;

-- name: DeleteUnit :one
DELETE FROM units WHERE id = $1
RETURNING id;
