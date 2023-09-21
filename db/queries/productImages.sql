-- name: CreateProductImage :one
INSERT INTO product_images (
    product_id,
    image,
    isdeleted,
    isblocked,
    isdefault
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id;

-- name: SelectProductImages :many
SELECT * FROM product_images;

-- name: UpdateProductImage :one
UPDATE product_images
SET
    product_id = $2,
    image = $3,
    isdeleted = $4,
    isblocked = $5,
    isdefault = $6
WHERE id = $1
RETURNING id;

-- name: DeleteProductImage :one
DELETE FROM product_images WHERE id = $1
RETURNING id;
