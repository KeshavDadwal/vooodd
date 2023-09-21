// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: offerCategories.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOfferCategory = `-- name: CreateOfferCategory :one
INSERT INTO offers_categories (deal_id, offer_id)
VALUES ($1, $2)
RETURNING id
`

type CreateOfferCategoryParams struct {
	DealID  pgtype.Int4 `json:"deal_id"`
	OfferID pgtype.Int4 `json:"offer_id"`
}

func (q *Queries) CreateOfferCategory(ctx context.Context, arg CreateOfferCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, createOfferCategory, arg.DealID, arg.OfferID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteOfferCategory = `-- name: DeleteOfferCategory :one
DELETE FROM offers_categories WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteOfferCategory(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteOfferCategory, id)
	err := row.Scan(&id)
	return id, err
}

const selectOfferCategories = `-- name: SelectOfferCategories :many
SELECT id, deal_id, offer_id, created_at FROM offers_categories
`

func (q *Queries) SelectOfferCategories(ctx context.Context) ([]OffersCategory, error) {
	rows, err := q.db.Query(ctx, selectOfferCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OffersCategory{}
	for rows.Next() {
		var i OffersCategory
		if err := rows.Scan(
			&i.ID,
			&i.DealID,
			&i.OfferID,
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

const updateOfferCategory = `-- name: UpdateOfferCategory :one
UPDATE offers_categories
SET
    deal_id = $2,
    offer_id = $3
WHERE id = $1
RETURNING id
`

type UpdateOfferCategoryParams struct {
	ID      int32       `json:"id"`
	DealID  pgtype.Int4 `json:"deal_id"`
	OfferID pgtype.Int4 `json:"offer_id"`
}

func (q *Queries) UpdateOfferCategory(ctx context.Context, arg UpdateOfferCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateOfferCategory, arg.ID, arg.DealID, arg.OfferID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
