-- name: CreateCoupon :one
INSERT INTO coupons (
    store_id,
    brand_id,
    product_id,
    name,
    code,
    description,
    path,
    isdeleted,
    isblocked,
    user_id,
    price,
    ispercent,
    start_date,
    end_date,
    unit_id,
    qty
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16
)
RETURNING id;

-- name: SelectCoupons :one
SELECT * FROM coupons WHERE id = $1;


-- name: SelectAllCoupons :many
SELECT * FROM coupons;

-- name: UpdateCoupon :one
UPDATE coupons
SET
    store_id = $2,
    brand_id = $3,
    product_id = $4,
    name = $5,
    code = $6,
    description = $7,
    path = $8,
    isdeleted = $9,
    isblocked = $10,
    user_id = $11,
    price = $12,
    ispercent = $13,
    start_date = $14,
    end_date = $15,
    unit_id = $16,
    qty = $17
WHERE id = $1
RETURNING id;

-- name: DeleteCoupon :one
DELETE FROM coupons WHERE id = $1
RETURNING id;

