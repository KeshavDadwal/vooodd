// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: offerViews.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOfferView = `-- name: CreateOfferView :one
INSERT INTO offer_views (
    product_id,
    offer_id,
    mall_id,
    store_id,
    store_location_id,
    user_id,
    viewed,
    view_date,
    is_clicked
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
    $9
)
RETURNING id
`

type CreateOfferViewParams struct {
	ProductID       int32       `json:"product_id"`
	OfferID         int32       `json:"offer_id"`
	MallID          int32       `json:"mall_id"`
	StoreID         int32       `json:"store_id"`
	StoreLocationID int32       `json:"store_location_id"`
	UserID          int32       `json:"user_id"`
	Viewed          pgtype.Int4 `json:"viewed"`
	ViewDate        pgtype.Date `json:"view_date"`
	IsClicked       pgtype.Bool `json:"is_clicked"`
}

func (q *Queries) CreateOfferView(ctx context.Context, arg CreateOfferViewParams) (int32, error) {
	row := q.db.QueryRow(ctx, createOfferView,
		arg.ProductID,
		arg.OfferID,
		arg.MallID,
		arg.StoreID,
		arg.StoreLocationID,
		arg.UserID,
		arg.Viewed,
		arg.ViewDate,
		arg.IsClicked,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteOfferView = `-- name: DeleteOfferView :one
DELETE FROM offer_views WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteOfferView(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteOfferView, id)
	err := row.Scan(&id)
	return id, err
}

const selectOfferViewsByOfferID = `-- name: SelectOfferViewsByOfferID :many
SELECT id, product_id, offer_id, mall_id, store_id, store_location_id, user_id, viewed, view_date, is_clicked, created_at FROM offer_views WHERE offer_id = $1
`

func (q *Queries) SelectOfferViewsByOfferID(ctx context.Context, offerID int32) ([]OfferView, error) {
	rows, err := q.db.Query(ctx, selectOfferViewsByOfferID, offerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OfferView{}
	for rows.Next() {
		var i OfferView
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.OfferID,
			&i.MallID,
			&i.StoreID,
			&i.StoreLocationID,
			&i.UserID,
			&i.Viewed,
			&i.ViewDate,
			&i.IsClicked,
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