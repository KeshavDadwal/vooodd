-- name: CreateBrand :one
INSERT INTO brands (name, logo, isdeleted, isblocked)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectBrands :one
SELECT * FROM brands WHERE id = $1;

-- name: SelectAllBrands :many
SELECT * FROM brands;

-- name: UpdateBrand :one
UPDATE brands
SET name = $2, logo = $3, isdeleted = $4, isblocked = $5
WHERE id = $1
RETURNING id;

-- name: DeleteBrand :one
DELETE FROM brands WHERE id = $1
RETURNING id;
