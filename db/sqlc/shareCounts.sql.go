// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: shareCounts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createShareCount = `-- name: CreateShareCount :one
INSERT INTO share_counts (
    language_id,
    product_category_id,
    share_type
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING id
`

type CreateShareCountParams struct {
	LanguageID        int32       `json:"language_id"`
	ProductCategoryID int32       `json:"product_category_id"`
	ShareType         pgtype.Text `json:"share_type"`
}

func (q *Queries) CreateShareCount(ctx context.Context, arg CreateShareCountParams) (int32, error) {
	row := q.db.QueryRow(ctx, createShareCount, arg.LanguageID, arg.ProductCategoryID, arg.ShareType)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteShareCount = `-- name: DeleteShareCount :one
DELETE FROM share_counts WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteShareCount(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteShareCount, id)
	err := row.Scan(&id)
	return id, err
}

const selectShareCounts = `-- name: SelectShareCounts :many
SELECT id, language_id, product_category_id, share_type, created_at FROM share_counts
`

func (q *Queries) SelectShareCounts(ctx context.Context) ([]ShareCount, error) {
	rows, err := q.db.Query(ctx, selectShareCounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ShareCount{}
	for rows.Next() {
		var i ShareCount
		if err := rows.Scan(
			&i.ID,
			&i.LanguageID,
			&i.ProductCategoryID,
			&i.ShareType,
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

const updateShareCount = `-- name: UpdateShareCount :one
UPDATE share_counts
SET
    language_id = $2,
    product_category_id = $3,
    share_type = $4
WHERE id = $1
RETURNING id
`

type UpdateShareCountParams struct {
	ID                int32       `json:"id"`
	LanguageID        int32       `json:"language_id"`
	ProductCategoryID int32       `json:"product_category_id"`
	ShareType         pgtype.Text `json:"share_type"`
}

func (q *Queries) UpdateShareCount(ctx context.Context, arg UpdateShareCountParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateShareCount,
		arg.ID,
		arg.LanguageID,
		arg.ProductCategoryID,
		arg.ShareType,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
