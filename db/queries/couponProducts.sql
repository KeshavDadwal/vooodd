-- name: CreateCouponProduct :one
INSERT INTO coupon_products (coupon_id, product_id)
VALUES ($1, $2)
RETURNING id;

-- name: SelectAllCouponProducts :many
SELECT * FROM coupon_products;

-- name: DeleteCouponProduct :one
DELETE FROM coupon_products WHERE coupon_id = $1 AND product_id = $2 RETURNING id;


