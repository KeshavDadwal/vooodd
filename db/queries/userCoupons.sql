-- name: CreateUserCoupon :one
INSERT INTO user_coupons (
    user_id,
    coupon_id
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectUserCoupons :many
SELECT * FROM user_coupons WHERE user_id = $1;

-- name: DeleteUserCoupon :one
DELETE FROM user_coupons WHERE id = $1
RETURNING id;
