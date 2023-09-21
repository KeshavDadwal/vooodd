-- name: CreateCompetitor :one
INSERT INTO competitors (name, logo, isdeleted, isblocked)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectCompetitors :one
SELECT * FROM competitors WHERE id = $1;

-- name: SelectAllCompetitors :many
SELECT * FROM competitors;

-- name: UpdateCompetitor :one
UPDATE competitors
SET name = $2, logo = $3, isdeleted = $4, isblocked = $5
WHERE id = $1
RETURNING id;

-- name: DeleteCompetitor :one
DELETE FROM competitors WHERE id = $1
RETURNING id;

