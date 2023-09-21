-- name: CreateStoreCategory :one
INSERT INTO store_categories (
    deal_id,
    store_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectStoreCategories :many
SELECT * FROM store_categories;

-- name: UpdateStoreCategory :one
UPDATE store_categories
SET
    deal_id = $2,
    store_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteStoreCategory :one
DELETE FROM store_categories WHERE id = $1
RETURNING id;

