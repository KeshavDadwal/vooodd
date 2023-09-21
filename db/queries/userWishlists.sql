-- name: CreateUserWishlistItem :one
INSERT INTO user_wishlists (
    user_id,
    product_id,
    store_location_id,
    wish_price,
    max_price,
    immediately,
    date_till
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

-- name: GetUserWishlistItems :many
SELECT
    id,
    user_id,
    product_id,
    store_location_id,
    wish_price,
    max_price,
    immediately,
    date_till,
    created_at
FROM
    user_wishlists
WHERE
    user_id = $1;

-- name: UpdateUserWishlistItem :one
UPDATE user_wishlists
SET
    product_id = $2,
    store_location_id = $3,
    wish_price = $4,
    max_price = $5,
    immediately = $6,
    date_till = $7
WHERE
    id = $1
    RETURNING id;

-- name: DeleteUserWishlistItem :one
DELETE FROM user_wishlists
WHERE
    id = $1
    RETURNING id;

