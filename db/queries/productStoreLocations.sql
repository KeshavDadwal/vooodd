-- name: CreateProductStoreLocation :one
INSERT INTO product_store_locations (
    product_id,
    store_location_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectProductStoreLocations :many
SELECT * FROM product_store_locations;

-- name: DeleteProductStoreLocation :one
DELETE FROM product_store_locations WHERE id = $1
RETURNING id;
