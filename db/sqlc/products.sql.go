// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: products.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    brand_id,
    product_category_id,
    name,
    sku,
    description,
    short_description,
    new_from,
    new_end,
    isdeleted,
    isblocked,
    isbestseller,
    isfeatured,
    meta_title,
    meta_keywords,
    meta_description,
    price,
    unit_id
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
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17
)
RETURNING id
`

type CreateProductParams struct {
	BrandID           pgtype.Int4 `json:"brand_id"`
	ProductCategoryID pgtype.Int4 `json:"product_category_id"`
	Name              string      `json:"name"`
	Sku               string      `json:"sku"`
	Description       pgtype.Text `json:"description"`
	ShortDescription  pgtype.Text `json:"short_description"`
	NewFrom           pgtype.Date `json:"new_from"`
	NewEnd            pgtype.Date `json:"new_end"`
	Isdeleted         pgtype.Bool `json:"isdeleted"`
	Isblocked         pgtype.Bool `json:"isblocked"`
	Isbestseller      pgtype.Bool `json:"isbestseller"`
	Isfeatured        pgtype.Bool `json:"isfeatured"`
	MetaTitle         string      `json:"meta_title"`
	MetaKeywords      string      `json:"meta_keywords"`
	MetaDescription   string      `json:"meta_description"`
	Price             int32       `json:"price"`
	UnitID            pgtype.Int4 `json:"unit_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (int32, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.BrandID,
		arg.ProductCategoryID,
		arg.Name,
		arg.Sku,
		arg.Description,
		arg.ShortDescription,
		arg.NewFrom,
		arg.NewEnd,
		arg.Isdeleted,
		arg.Isblocked,
		arg.Isbestseller,
		arg.Isfeatured,
		arg.MetaTitle,
		arg.MetaKeywords,
		arg.MetaDescription,
		arg.Price,
		arg.UnitID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteProduct = `-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteProduct, id)
	err := row.Scan(&id)
	return id, err
}

const selectProducts = `-- name: SelectProducts :many
SELECT id, brand_id, product_category_id, name, sku, description, short_description, new_from, new_end, isdeleted, isblocked, isbestseller, isfeatured, meta_title, meta_keywords, meta_description, price, unit_id, created_at FROM products WHERE isdeleted = false AND isblocked = false
`

func (q *Queries) SelectProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, selectProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.BrandID,
			&i.ProductCategoryID,
			&i.Name,
			&i.Sku,
			&i.Description,
			&i.ShortDescription,
			&i.NewFrom,
			&i.NewEnd,
			&i.Isdeleted,
			&i.Isblocked,
			&i.Isbestseller,
			&i.Isfeatured,
			&i.MetaTitle,
			&i.MetaKeywords,
			&i.MetaDescription,
			&i.Price,
			&i.UnitID,
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

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET
    brand_id = $2,
    product_category_id = $3,
    name = $4,
    sku = $5,
    description = $6,
    short_description = $7,
    new_from = $8,
    new_end = $9,
    isdeleted = $10,
    isblocked = $11,
    isbestseller = $12,
    isfeatured = $13,
    meta_title = $14,
    meta_keywords = $15,
    meta_description = $16,
    price = $17,
    unit_id = $18
WHERE id = $1
RETURNING id
`

type UpdateProductParams struct {
	ID                int32       `json:"id"`
	BrandID           pgtype.Int4 `json:"brand_id"`
	ProductCategoryID pgtype.Int4 `json:"product_category_id"`
	Name              string      `json:"name"`
	Sku               string      `json:"sku"`
	Description       pgtype.Text `json:"description"`
	ShortDescription  pgtype.Text `json:"short_description"`
	NewFrom           pgtype.Date `json:"new_from"`
	NewEnd            pgtype.Date `json:"new_end"`
	Isdeleted         pgtype.Bool `json:"isdeleted"`
	Isblocked         pgtype.Bool `json:"isblocked"`
	Isbestseller      pgtype.Bool `json:"isbestseller"`
	Isfeatured        pgtype.Bool `json:"isfeatured"`
	MetaTitle         string      `json:"meta_title"`
	MetaKeywords      string      `json:"meta_keywords"`
	MetaDescription   string      `json:"meta_description"`
	Price             int32       `json:"price"`
	UnitID            pgtype.Int4 `json:"unit_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.BrandID,
		arg.ProductCategoryID,
		arg.Name,
		arg.Sku,
		arg.Description,
		arg.ShortDescription,
		arg.NewFrom,
		arg.NewEnd,
		arg.Isdeleted,
		arg.Isblocked,
		arg.Isbestseller,
		arg.Isfeatured,
		arg.MetaTitle,
		arg.MetaKeywords,
		arg.MetaDescription,
		arg.Price,
		arg.UnitID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}