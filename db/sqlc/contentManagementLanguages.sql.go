// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: contentManagementLanguages.sql

package db

import (
	"context"
)

const createContentManagementLanguage = `-- name: CreateContentManagementLanguage :one
INSERT INTO content_management_languages (content_management_id, language_id, name, title, content)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateContentManagementLanguageParams struct {
	ContentManagementID int32  `json:"content_management_id"`
	LanguageID          int32  `json:"language_id"`
	Name                string `json:"name"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

func (q *Queries) CreateContentManagementLanguage(ctx context.Context, arg CreateContentManagementLanguageParams) (int32, error) {
	row := q.db.QueryRow(ctx, createContentManagementLanguage,
		arg.ContentManagementID,
		arg.LanguageID,
		arg.Name,
		arg.Title,
		arg.Content,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteContentManagementLanguage = `-- name: DeleteContentManagementLanguage :one
DELETE FROM content_management_languages WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteContentManagementLanguage(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteContentManagementLanguage, id)
	err := row.Scan(&id)
	return id, err
}

const selectAllContentManagementLanguages = `-- name: SelectAllContentManagementLanguages :many
SELECT id, content_management_id, language_id, name, title, content, created_at FROM content_management_languages
`

func (q *Queries) SelectAllContentManagementLanguages(ctx context.Context) ([]ContentManagementLanguage, error) {
	rows, err := q.db.Query(ctx, selectAllContentManagementLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ContentManagementLanguage{}
	for rows.Next() {
		var i ContentManagementLanguage
		if err := rows.Scan(
			&i.ID,
			&i.ContentManagementID,
			&i.LanguageID,
			&i.Name,
			&i.Title,
			&i.Content,
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

const selectContentManagementLanguages = `-- name: SelectContentManagementLanguages :one
SELECT id, content_management_id, language_id, name, title, content, created_at FROM content_management_languages WHERE id = $1
`

func (q *Queries) SelectContentManagementLanguages(ctx context.Context, id int32) (ContentManagementLanguage, error) {
	row := q.db.QueryRow(ctx, selectContentManagementLanguages, id)
	var i ContentManagementLanguage
	err := row.Scan(
		&i.ID,
		&i.ContentManagementID,
		&i.LanguageID,
		&i.Name,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const updateContentManagementLanguage = `-- name: UpdateContentManagementLanguage :one
UPDATE content_management_languages
SET name = $3, title = $4, content = $5
WHERE id = $1 AND content_management_id = $2
RETURNING id
`

type UpdateContentManagementLanguageParams struct {
	ID                  int32  `json:"id"`
	ContentManagementID int32  `json:"content_management_id"`
	Name                string `json:"name"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

func (q *Queries) UpdateContentManagementLanguage(ctx context.Context, arg UpdateContentManagementLanguageParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateContentManagementLanguage,
		arg.ID,
		arg.ContentManagementID,
		arg.Name,
		arg.Title,
		arg.Content,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
