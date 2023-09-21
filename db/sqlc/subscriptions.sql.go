// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: subscriptions.sql

package db

import (
	"context"
)

const createSubscription = `-- name: CreateSubscription :one
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
RETURNING id
`

type CreateSubscriptionParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	MallID    int32  `json:"mall_id"`
	Zip       string `json:"zip"`
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) (int32, error) {
	row := q.db.QueryRow(ctx, createSubscription,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.MallID,
		arg.Zip,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteSubscription = `-- name: DeleteSubscription :one
DELETE FROM subscriptions WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteSubscription(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteSubscription, id)
	err := row.Scan(&id)
	return id, err
}

const selectSubscriptions = `-- name: SelectSubscriptions :many
SELECT id, first_name, last_name, email, mall_id, zip, created_at FROM subscriptions
`

func (q *Queries) SelectSubscriptions(ctx context.Context) ([]Subscription, error) {
	rows, err := q.db.Query(ctx, selectSubscriptions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Subscription{}
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.MallID,
			&i.Zip,
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

const updateSubscription = `-- name: UpdateSubscription :one
UPDATE subscriptions
SET
    first_name = $2,
    last_name = $3,
    email = $4,
    mall_id = $5,
    zip = $6
WHERE id = $1
RETURNING id, first_name, last_name, email, mall_id, zip, created_at
`

type UpdateSubscriptionParams struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	MallID    int32  `json:"mall_id"`
	Zip       string `json:"zip"`
}

func (q *Queries) UpdateSubscription(ctx context.Context, arg UpdateSubscriptionParams) (Subscription, error) {
	row := q.db.QueryRow(ctx, updateSubscription,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.MallID,
		arg.Zip,
	)
	var i Subscription
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.MallID,
		&i.Zip,
		&i.CreatedAt,
	)
	return i, err
}
