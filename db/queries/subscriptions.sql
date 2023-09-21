-- name: CreateSubscription :one
INSERT INTO subscriptions (
    first_name,
    last_name,
    email,
    mall_id,
    zip
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id;

-- name: SelectSubscriptions :many
SELECT * FROM subscriptions;

-- name: UpdateSubscription :one
UPDATE subscriptions
SET
    first_name = $2,
    last_name = $3,
    email = $4,
    mall_id = $5,
    zip = $6
WHERE id = $1
RETURNING *;

-- name: DeleteSubscription :one
DELETE FROM subscriptions WHERE id = $1
RETURNING id;
