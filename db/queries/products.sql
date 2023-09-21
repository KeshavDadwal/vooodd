-- name: CreateProduct :one
INSERT INTO products (
    brand_id,
    product_category_id,
    name,
    sku,
    description,
    short_description,
    new_from,
    new_end,
    isdeleted,
    isblocked,
    isbestseller,
    isfeatured,
    meta_title,
    meta_keywords,
    meta_description,
    price,
    unit_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17
)
RETURNING id;

-- name: SelectProducts :many
SELECT * FROM products WHERE isdeleted = false AND isblocked = false;

-- name: UpdateProduct :one
UPDATE products
SET
    brand_id = $2,
    product_category_id = $3,
    name = $4,
    sku = $5,
    description = $6,
    short_description = $7,
    new_from = $8,
    new_end = $9,
    isdeleted = $10,
    isblocked = $11,
    isbestseller = $12,
    isfeatured = $13,
    meta_title = $14,
    meta_keywords = $15,
    meta_description = $16,
    price = $17,
    unit_id = $18
WHERE id = $1
RETURNING id;

-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1
RETURNING id;


