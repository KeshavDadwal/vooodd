-- name: CreateUserShoppingListItem :one
INSERT INTO user_shoping_lists (
    user_shoping_list_name_id,
    product_id,
    user_id,
    store_location_id,
    qty
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id;

-- name: GetUserShoppingList :many
SELECT
    id,
    user_shoping_list_name_id,
    product_id,
    user_id,
    store_location_id,
    qty,
    created_at
FROM
    user_shoping_lists
WHERE
    user_id = $1;

-- name: UpdateUserShoppingListItem :one
UPDATE user_shoping_lists
SET
    user_shoping_list_name_id = $2,
    product_id = $3,
    store_location_id = $4,
    qty = $5
WHERE
    id = $1
    RETURNING id;

-- name: DeleteUserShoppingListItem :one
DELETE FROM user_shoping_lists
WHERE
    id = $1
    RETURNING id;
