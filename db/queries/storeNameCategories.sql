-- name: CreateStoreNameCategory :one
INSERT INTO store_name_categories (
    parent_id,
    lft,
    rght,
    name
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: SelectStoreNameCategories :many
SELECT * FROM store_name_categories;

-- name: UpdateStoreNameCategory :one
UPDATE store_name_categories
SET
    parent_id = $2,
    lft = $3,
    rght = $4,
    name = $5
WHERE id = $1
RETURNING id;

-- name: DeleteStoreNameCategory :one
DELETE FROM store_name_categories WHERE id = $1
RETURNING id;


