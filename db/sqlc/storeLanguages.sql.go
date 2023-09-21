// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: storeLanguages.sql

package db

import (
	"context"
)

const createStoreLanguage = `-- name: CreateStoreLanguage :one
INSERT INTO store_languages (
    language_id,
    store_id,
    name,
    description
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id
`

type CreateStoreLanguageParams struct {
	LanguageID  int32  `json:"language_id"`
	StoreID     int32  `json:"store_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateStoreLanguage(ctx context.Context, arg CreateStoreLanguageParams) (int32, error) {
	row := q.db.QueryRow(ctx, createStoreLanguage,
		arg.LanguageID,
		arg.StoreID,
		arg.Name,
		arg.Description,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteStoreLanguage = `-- name: DeleteStoreLanguage :one
DELETE FROM store_languages WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteStoreLanguage(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteStoreLanguage, id)
	err := row.Scan(&id)
	return id, err
}

const selectStoreLanguages = `-- name: SelectStoreLanguages :many
SELECT id, language_id, store_id, name, description, created_at FROM store_languages
`

func (q *Queries) SelectStoreLanguages(ctx context.Context) ([]StoreLanguage, error) {
	rows, err := q.db.Query(ctx, selectStoreLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StoreLanguage{}
	for rows.Next() {
		var i StoreLanguage
		if err := rows.Scan(
			&i.ID,
			&i.LanguageID,
			&i.StoreID,
			&i.Name,
			&i.Description,
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

const updateStoreLanguage = `-- name: UpdateStoreLanguage :one
UPDATE store_languages
SET
    language_id = $2,
    store_id = $3,
    name = $4,
    description = $5
WHERE id = $1
RETURNING id
`

type UpdateStoreLanguageParams struct {
	ID          int32  `json:"id"`
	LanguageID  int32  `json:"language_id"`
	StoreID     int32  `json:"store_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) UpdateStoreLanguage(ctx context.Context, arg UpdateStoreLanguageParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateStoreLanguage,
		arg.ID,
		arg.LanguageID,
		arg.StoreID,
		arg.Name,
		arg.Description,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
