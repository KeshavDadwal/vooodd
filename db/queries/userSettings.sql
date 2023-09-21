-- name: CreateUserSetting :one
INSERT INTO user_settings (
    user_id,
    mall_radius,
    store_radius
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING id;

-- name: GetUserSettingsByUserId :one
SELECT
    id,
    user_id,
    mall_radius,
    store_radius,
    created_at
FROM
    user_settings
WHERE
    user_id = $1;

-- name: UpdateUserSetting :one
UPDATE user_settings
SET
    mall_radius = $2,
    store_radius = $3
WHERE
    user_id = $1
    RETURNING id;

-- name: DeleteUserSetting :one
DELETE FROM user_settings
WHERE
    user_id = $1
    RETURNING id;


