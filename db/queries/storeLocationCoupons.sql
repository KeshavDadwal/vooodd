-- name: CreateStoreLocationCoupon :one
INSERT INTO store_location_coupons (
    store_location_id,
    product_id,
    coupon_id,
    price,
    coupon_price,
    is_percent,
    start_date,
    end_date
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

-- name: SelectStoreLocationCoupons :many
SELECT * FROM store_location_coupons;

-- name: UpdateStoreLocationCoupon :one
UPDATE store_location_coupons
SET
    store_location_id = $2,
    product_id = $3,
    coupon_id = $4,
    price = $5,
    coupon_price = $6,
    is_percent = $7,
    start_date = $8,
    end_date = $9
WHERE id = $1
RETURNING *;

-- name: DeleteStoreLocationCoupon :one
DELETE FROM store_location_coupons WHERE id = $1
RETURNING id;
