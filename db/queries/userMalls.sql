-- name: CreateUserMallMapping :one
INSERT INTO user_malls (
    user_id,
    mall_id,
    device_id
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING id;

-- name: GetUserMallMappingByUserId :one
SELECT
    id,
    user_id,
    mall_id,
    device_id,
    created_at
FROM
    user_malls
WHERE
    user_id = $1;

-- name: UpdateUserMallMapping :one
UPDATE user_malls
SET
    mall_id = $2
WHERE
    user_id = $1
    RETURNING id;


-- name: DeleteUserMallMappingByUserId :one
DELETE FROM user_malls
WHERE
    user_id = $1
    RETURNING id;
