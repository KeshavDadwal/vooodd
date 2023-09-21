-- name: CreateProductTest :one
INSERT INTO product_tests (
    brand_id,
    product_category_id,
    name,
    sku,
    description,
    short_description,
    weight,
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
    $17,
    $18
)
RETURNING id;

-- name: SelectProductTests :many
SELECT * FROM product_tests;

-- name: UpdateProductTest :one
UPDATE product_tests
SET
    brand_id = $2,
    product_category_id = $3,
    name = $4,
    sku = $5,
    description = $6,
    short_description = $7,
    weight = $8,
    new_from = $9,
    new_end = $10,
    isdeleted = $11,
    isblocked = $12,
    isbestseller = $13,
    isfeatured = $14,
    meta_title = $15,
    meta_keywords = $16,
    meta_description = $17,
    price = $18,
    unit_id = $19
WHERE id = $1
RETURNING id;

-- name: DeleteProductTest :one
DELETE FROM product_tests WHERE id = $1
RETURNING id;

