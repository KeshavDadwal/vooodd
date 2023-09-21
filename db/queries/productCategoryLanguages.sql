-- name: CreateProductCategoryLanguage :one
INSERT INTO product_category_languages (
    language_id,
    product_category_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectProductCategoryLanguages :many
SELECT * FROM product_category_languages;

-- name: UpdateProductCategoryLanguage :one
UPDATE product_category_languages
SET
    language_id = $2,
    product_category_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteProductCategoryLanguage :one
DELETE FROM product_category_languages WHERE id = $1
RETURNING id;

