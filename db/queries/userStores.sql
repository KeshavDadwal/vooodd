-- name: CreateUserStoreRelationship :one
INSERT INTO user_stores (
    user_id,
    store_location_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: GetUserFavoriteStores :many
SELECT
    id,
    user_id,
    store_location_id,
    created_at
FROM
    user_stores
WHERE
    user_id = $1;

-- name: UpdateUserStoreRelationship :one
UPDATE user_stores
SET
    store_location_id = $2
WHERE
    id = $1
    RETURNING id;

-- name: DeleteUserStoreRelationship :one
DELETE FROM user_stores
WHERE
    id = $1
    RETURNING id;

