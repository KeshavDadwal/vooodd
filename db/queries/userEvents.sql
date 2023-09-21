-- name: CreateUserEvent :one
INSERT INTO user_events (
    user_id,
    name,
    date
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING id;

-- name: SelectUserEvents :many
SELECT * FROM user_events WHERE user_id = $1;

-- name: UpdateUserEvent :one
UPDATE user_events
SET
    name = $2,
    date = $3
WHERE id = $1
RETURNING id;

-- name: DeleteUserEvent :one
DELETE FROM user_events WHERE id = $1
RETURNING id;
