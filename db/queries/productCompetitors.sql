-- name: CreateProductCompetitor :one
INSERT INTO product_competitors (
    product_id,
    competitor_id,
    store_location_id,
    price
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: GetProductCompetitors :many
SELECT
    id,
    product_id,
    competitor_id,
    store_location_id,
    price,
    created_at
FROM
    product_competitors
WHERE
    product_id = $1;

-- name: UpdateProductCompetitorPrice :one
UPDATE product_competitors
SET
    price = $2
WHERE
    id = $1
    RETURNING *;

-- name: DeleteProductCompetitor :one
DELETE FROM product_competitors
WHERE
    id = $1
    RETURNING id;

