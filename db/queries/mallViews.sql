-- name: CreateMallView :one
INSERT INTO mall_views (
    user_id,
    mall_id,
    store_id,
    visited,
    device_id,
    from_device,
    is_physical
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING id;

-- name: SelectMallViews :many
SELECT * FROM mall_views;

-- name: UpdateMallView :one
UPDATE mall_views
SET
    user_id = $2,
    mall_id = $3,
    store_id = $4,
    visited = $5,
    device_id = $6,
    from_device = $7,
    is_physical = $8
WHERE id = $1
RETURNING *;


-- name: DeleteMallView :one
DELETE FROM mall_views WHERE id = $1
RETURNING id;
