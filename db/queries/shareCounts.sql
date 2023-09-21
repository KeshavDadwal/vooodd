-- name: CreateShareCount :one
INSERT INTO share_counts (
    language_id,
    product_category_id,
    share_type
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING id;

-- name: SelectShareCounts :many
SELECT * FROM share_counts;


-- name: UpdateShareCount :one
UPDATE share_counts
SET
    language_id = $2,
    product_category_id = $3,
    share_type = $4
WHERE id = $1
RETURNING id;

-- name: DeleteShareCount :one
DELETE FROM share_counts WHERE id = $1
RETURNING id;
