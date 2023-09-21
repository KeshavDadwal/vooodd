-- name: CreateUserMallPermission :one
INSERT INTO user_mall_permissions (
    user_id,
    mall_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: GetUserMallPermissionsByUserId :one
SELECT
    id,
    user_id,
    mall_id,
    created_at
FROM
    user_mall_permissions
WHERE
    user_id = $1;

-- name: UpdateUserMallPermission :one
UPDATE user_mall_permissions
SET
    mall_id = $2
WHERE
    user_id = $1
    RETURNING id;

-- name: DeleteUserMallPermission :one
DELETE FROM user_mall_permissions
WHERE
    user_id = $1
    AND mall_id = $2
    RETURNING id;

