// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: productCategories.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProductCategory = `-- name: CreateProductCategory :one
INSERT INTO product_categories (parent_id, lft, rght, name, logo, store_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`

type CreateProductCategoryParams struct {
	ParentID pgtype.Int4 `json:"parent_id"`
	Lft      pgtype.Int4 `json:"lft"`
	Rght     pgtype.Int4 `json:"rght"`
	Name     pgtype.Text `json:"name"`
	Logo     pgtype.Text `json:"logo"`
	StoreID  pgtype.Int4 `json:"store_id"`
}

func (q *Queries) CreateProductCategory(ctx context.Context, arg CreateProductCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, createProductCategory,
		arg.ParentID,
		arg.Lft,
		arg.Rght,
		arg.Name,
		arg.Logo,
		arg.StoreID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteProductCategory = `-- name: DeleteProductCategory :one
DELETE FROM product_categories WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteProductCategory(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteProductCategory, id)
	err := row.Scan(&id)
	return id, err
}

const selectProductCategories = `-- name: SelectProductCategories :many
SELECT id, parent_id, lft, rght, name, logo, store_id, created_at FROM product_categories
`

func (q *Queries) SelectProductCategories(ctx context.Context) ([]ProductCategory, error) {
	rows, err := q.db.Query(ctx, selectProductCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductCategory{}
	for rows.Next() {
		var i ProductCategory
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Lft,
			&i.Rght,
			&i.Name,
			&i.Logo,
			&i.StoreID,
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

const updateProductCategory = `-- name: UpdateProductCategory :one
UPDATE product_categories
SET parent_id = $2, lft = $3, rght = $4, name = $5, logo = $6, store_id = $7
WHERE id = $1
RETURNING id
`

type UpdateProductCategoryParams struct {
	ID       int32       `json:"id"`
	ParentID pgtype.Int4 `json:"parent_id"`
	Lft      pgtype.Int4 `json:"lft"`
	Rght     pgtype.Int4 `json:"rght"`
	Name     pgtype.Text `json:"name"`
	Logo     pgtype.Text `json:"logo"`
	StoreID  pgtype.Int4 `json:"store_id"`
}

func (q *Queries) UpdateProductCategory(ctx context.Context, arg UpdateProductCategoryParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateProductCategory,
		arg.ID,
		arg.ParentID,
		arg.Lft,
		arg.Rght,
		arg.Name,
		arg.Logo,
		arg.StoreID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
