-- name: CreateDislikeCount :one
INSERT INTO dis_like_counts (
    user_id,
    offer_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectAllDislikeCounts :many
SELECT * FROM dis_like_counts;

-- name: UpdateDislikeCount :one
UPDATE dis_like_counts
SET
    user_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteDislikeCount :one
DELETE FROM dis_like_counts WHERE id = $1
RETURNING id;
