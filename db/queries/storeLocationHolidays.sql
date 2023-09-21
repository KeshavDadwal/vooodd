-- name: CreateStoreLocationHoliday :one
INSERT INTO store_location_holidays (
    store_location_id,
    name,
    start_date,
    end_date
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: SelectStoreLocationHolidays :many
SELECT * FROM store_location_holidays;

-- name: UpdateStoreLocationHoliday :one
UPDATE store_location_holidays
SET
    store_location_id = $2,
    name = $3,
    start_date = $4,
    end_date = $5
WHERE id = $1
RETURNING *;

-- name: DeleteStoreLocationHoliday :one
DELETE FROM store_location_holidays WHERE id = $1
RETURNING id;
