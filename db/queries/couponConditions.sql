-- name: CreateCouponCondition :one
INSERT INTO coupon_conditions (
    coupon_id,
    idx,
    type,
    type_id
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: SelectCouponConditions :one
SELECT * FROM coupon_conditions WHERE id = $1;


-- name: SelectAllCouponConditions :many
SELECT * FROM coupon_conditions;


-- name: UpdateCouponCondition :one
UPDATE coupon_conditions
SET
    idx = $2,
    type = $3,
    type_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCouponCondition :one
DELETE FROM coupon_conditions WHERE id = $1
RETURNING id;
