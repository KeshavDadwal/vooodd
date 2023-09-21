-- name: CreateUserFavoriteStore :one
INSERT INTO user_favourite_stores (
    user_id,
    store_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectUserFavoriteStores :many
SELECT * FROM user_favourite_stores WHERE user_id = $1;

-- name: UpdateUserFavoriteStore :one
UPDATE user_favourite_stores
SET
    store_id = $2
WHERE id = $1
RETURNING id;

-- name: DeleteUserFavoriteStore :one
DELETE FROM user_favourite_stores WHERE id = $1
RETURNING id;

