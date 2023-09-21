-- name: CreateUserAction :one
INSERT INTO user_actions (
    user_id,
    save_type,
    tablename,
    autoincrement_id,
    fieldname,
    value,
    old_value,
    note
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING id;

-- name: GetUserActionsByUserId :many
SELECT
    id,
    user_id,
    save_type,
    tablename,
    autoincrement_id,
    fieldname,
    value,
    old_value,
    note,
    created_at
FROM
    user_actions
WHERE
    user_id = $1;

-- name: UpdateUserAction :one
UPDATE user_actions
SET
    save_type = $2,
    tablename = $3,
    autoincrement_id = $4,
    fieldname = $5,
    value = $6,
    old_value = $7,
    note = $8
WHERE
    user_id = $1
    RETURNING id;


-- name: DeleteUserActionsByUserId :one
DELETE FROM user_actions
WHERE
    user_id = $1
    RETURNING id;

