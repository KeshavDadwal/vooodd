-- name: CreateStoreBusinessHours :one
INSERT INTO store_business_hours (
    store_location_id,
    day,
    start_time,
    end_time
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: SelectStoreBusinessHours :many
SELECT * FROM store_business_hours;

-- name: UpdateStoreBusinessHours :one
UPDATE store_business_hours
SET
    store_location_id = $2,
    day = $3,
    start_time = $4,
    end_time = $5
WHERE id = $1
RETURNING *;

-- name: DeleteStoreBusinessHours :one
DELETE FROM store_business_hours WHERE id = $1
RETURNING id;
