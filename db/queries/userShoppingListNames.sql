-- name: CreateUserShoppingListName :one
INSERT INTO user_shopping_list_names (
    user_id,
    name,
    how_obtain,
    frequency,
    type,
    isnotify
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id;

-- name: SelectUserShoppingListNames :many
SELECT * FROM user_shopping_list_names WHERE user_id = $1;

-- name: UpdateUserShoppingListName :one
UPDATE user_shopping_list_names
SET
    name = $2,
    how_obtain = $3,
    frequency = $4,
    type = $5,
    isnotify = $6,
    modified = NOW()
WHERE id = $1
RETURNING id;

-- name: DeleteUserShoppingListName :one
DELETE FROM user_shopping_list_names WHERE id = $1
RETURNING id;
