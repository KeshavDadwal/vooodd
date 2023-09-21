-- name: CreateUserLog :one
INSERT INTO user_logs (
    user_id,
    mall_id,
    store_id,
    offer_id,
    name,
    text
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

-- name: GetUserLogsByUserId :one
SELECT
    id,
    user_id,
    mall_id,
    store_id,
    offer_id,
    name,
    text,
    isblocked,
    isdeleted,
    created_at
FROM
    user_logs
WHERE
    user_id = $1;


-- name: UpdateUserLog :one
UPDATE user_logs
SET
    mall_id = $2,
    store_id = $3,
    offer_id = $4,
    name = $5,
    text = $6,
    isblocked = $7,
    isdeleted = $8
WHERE
    user_id = $1
    RETURNING id;

-- name: DeleteUserLogByUserId :one
DELETE FROM user_logs
WHERE
    user_id = $1
    RETURNING id;
