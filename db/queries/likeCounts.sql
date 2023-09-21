-- name: CreateLikeCount :one
INSERT INTO like_counts (
    user_id,
    offer_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectAlllikeCounts :one
SELECT * FROM like_counts;

-- name: DeleteLikeCount :one
DELETE FROM like_counts WHERE id = $1
RETURNING id;
