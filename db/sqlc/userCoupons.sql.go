// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: userCoupons.sql

package db

import (
	"context"
)

const createUserCoupon = `-- name: CreateUserCoupon :one
INSERT INTO user_coupons (
    user_id,
    coupon_id
)
VALUES (
    $1,
    $2
)
RETURNING id
`

type CreateUserCouponParams struct {
	UserID   int32 `json:"user_id"`
	CouponID int32 `json:"coupon_id"`
}

func (q *Queries) CreateUserCoupon(ctx context.Context, arg CreateUserCouponParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUserCoupon, arg.UserID, arg.CouponID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserCoupon = `-- name: DeleteUserCoupon :one
DELETE FROM user_coupons WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUserCoupon(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserCoupon, id)
	err := row.Scan(&id)
	return id, err
}

const selectUserCoupons = `-- name: SelectUserCoupons :many
SELECT id, user_id, coupon_id, created_at FROM user_coupons WHERE user_id = $1
`

func (q *Queries) SelectUserCoupons(ctx context.Context, userID int32) ([]UserCoupon, error) {
	rows, err := q.db.Query(ctx, selectUserCoupons, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserCoupon{}
	for rows.Next() {
		var i UserCoupon
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CouponID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
