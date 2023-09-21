// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: productCompetitors.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProductCompetitor = `-- name: CreateProductCompetitor :one
INSERT INTO product_competitors (
    product_id,
    competitor_id,
    store_location_id,
    price
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id
`

type CreateProductCompetitorParams struct {
	ProductID       int32          `json:"product_id"`
	CompetitorID    int32          `json:"competitor_id"`
	StoreLocationID int32          `json:"store_location_id"`
	Price           pgtype.Numeric `json:"price"`
}

func (q *Queries) CreateProductCompetitor(ctx context.Context, arg CreateProductCompetitorParams) (int32, error) {
	row := q.db.QueryRow(ctx, createProductCompetitor,
		arg.ProductID,
		arg.CompetitorID,
		arg.StoreLocationID,
		arg.Price,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteProductCompetitor = `-- name: DeleteProductCompetitor :one
DELETE FROM product_competitors
WHERE
    id = $1
    RETURNING id
`

func (q *Queries) DeleteProductCompetitor(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteProductCompetitor, id)
	err := row.Scan(&id)
	return id, err
}

const getProductCompetitors = `-- name: GetProductCompetitors :many
SELECT
    id,
    product_id,
    competitor_id,
    store_location_id,
    price,
    created_at
FROM
    product_competitors
WHERE
    product_id = $1
`

func (q *Queries) GetProductCompetitors(ctx context.Context, productID int32) ([]ProductCompetitor, error) {
	rows, err := q.db.Query(ctx, getProductCompetitors, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductCompetitor{}
	for rows.Next() {
		var i ProductCompetitor
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.CompetitorID,
			&i.StoreLocationID,
			&i.Price,
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

const updateProductCompetitorPrice = `-- name: UpdateProductCompetitorPrice :one
UPDATE product_competitors
SET
    price = $2
WHERE
    id = $1
    RETURNING id, product_id, competitor_id, store_location_id, price, created_at
`

type UpdateProductCompetitorPriceParams struct {
	ID    int32          `json:"id"`
	Price pgtype.Numeric `json:"price"`
}

func (q *Queries) UpdateProductCompetitorPrice(ctx context.Context, arg UpdateProductCompetitorPriceParams) (ProductCompetitor, error) {
	row := q.db.QueryRow(ctx, updateProductCompetitorPrice, arg.ID, arg.Price)
	var i ProductCompetitor
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.CompetitorID,
		&i.StoreLocationID,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}