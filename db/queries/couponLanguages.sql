-- name: CreateCouponLanguage :one
INSERT INTO coupon_languages (coupon_id, language_id, name, description, disclaimer)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: SelectAllCouponLanguages :many
SELECT * FROM coupon_languages;

-- name: DeleteCouponLanguage :one
DELETE FROM coupon_languages WHERE id = $1
RETURNING id;
