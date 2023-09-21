// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: deals.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDeal = `-- name: CreateDeal :one
INSERT INTO deals (parent_id, lft, rght, name)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateDealParams struct {
	ParentID pgtype.Int4 `json:"parent_id"`
	Lft      pgtype.Int4 `json:"lft"`
	Rght     pgtype.Int4 `json:"rght"`
	Name     pgtype.Text `json:"name"`
}

func (q *Queries) CreateDeal(ctx context.Context, arg CreateDealParams) (int32, error) {
	row := q.db.QueryRow(ctx, createDeal,
		arg.ParentID,
		arg.Lft,
		arg.Rght,
		arg.Name,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteDeal = `-- name: DeleteDeal :one
DELETE FROM deals WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteDeal(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteDeal, id)
	err := row.Scan(&id)
	return id, err
}

const selectAllDeals = `-- name: SelectAllDeals :many
SELECT id, parent_id, lft, rght, name, created_at FROM deals
`

func (q *Queries) SelectAllDeals(ctx context.Context) ([]Deal, error) {
	rows, err := q.db.Query(ctx, selectAllDeals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Deal{}
	for rows.Next() {
		var i Deal
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Lft,
			&i.Rght,
			&i.Name,
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

const selectDeals = `-- name: SelectDeals :one
SELECT id, parent_id, lft, rght, name, created_at FROM deals where id = $1
`

func (q *Queries) SelectDeals(ctx context.Context, id int32) (Deal, error) {
	row := q.db.QueryRow(ctx, selectDeals, id)
	var i Deal
	err := row.Scan(
		&i.ID,
		&i.ParentID,
		&i.Lft,
		&i.Rght,
		&i.Name,
		&i.CreatedAt,
	)
	return i, err
}

const updateDeal = `-- name: UpdateDeal :one
UPDATE deals
SET
    parent_id = $2,
    lft = $3,
    rght = $4,
    name = $5
WHERE id = $1
RETURNING id
`

type UpdateDealParams struct {
	ID       int32       `json:"id"`
	ParentID pgtype.Int4 `json:"parent_id"`
	Lft      pgtype.Int4 `json:"lft"`
	Rght     pgtype.Int4 `json:"rght"`
	Name     pgtype.Text `json:"name"`
}

func (q *Queries) UpdateDeal(ctx context.Context, arg UpdateDealParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateDeal,
		arg.ID,
		arg.ParentID,
		arg.Lft,
		arg.Rght,
		arg.Name,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
