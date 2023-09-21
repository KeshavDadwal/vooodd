-- name: CreateProductCategoryMap :one
INSERT INTO product_category_maps (
    product_id,
    product_category_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectProductCategoryMaps :many
SELECT * FROM product_category_maps;

-- name: UpdateProductCategoryMap :one
UPDATE product_category_maps
SET
    product_id = $2,
    product_category_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteProductCategoryMap :one
DELETE FROM product_category_maps WHERE id = $1
RETURNING id;
