-- name: CreateProductLanguage :one
INSERT INTO product_languages (
    product_id,
    language_id,
    name,
    sku,
    description,
    short_description,
    weight,
    meta_title,
    meta_keywords,
    meta_description
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
    $10
)
RETURNING id;

-- name: SelectProductLanguages :many
SELECT * FROM product_languages;

-- name: UpdateProductLanguage :one
UPDATE product_languages
SET
    product_id = $2,
    language_id = $3,
    name = $4,
    sku = $5,
    description = $6,
    short_description = $7,
    weight = $8,
    meta_title = $9,
    meta_keywords = $10,
    meta_description = $11
WHERE id = $1
RETURNING id;

-- name: DeleteProductLanguage :one
DELETE FROM product_languages WHERE id = $1
RETURNING id;

