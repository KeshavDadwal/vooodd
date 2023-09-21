-- name: CreateProductCategory :one
INSERT INTO product_categories (parent_id, lft, rght, name, logo, store_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: SelectProductCategories :many
SELECT * FROM product_categories;

-- name: UpdateProductCategory :one
UPDATE product_categories
SET parent_id = $2, lft = $3, rght = $4, name = $5, logo = $6, store_id = $7
WHERE id = $1
RETURNING id;

-- name: DeleteProductCategory :one
DELETE FROM product_categories WHERE id = $1
RETURNING id;
