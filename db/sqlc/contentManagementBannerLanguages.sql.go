// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: contentManagementBannerLanguages.sql

package db

import (
	"context"
)

const createContentManagementBannerLanguage = `-- name: CreateContentManagementBannerLanguage :one
INSERT INTO content_management_banner_languages (content_management_id, language_id, name, title, content)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateContentManagementBannerLanguageParams struct {
	ContentManagementID int32  `json:"content_management_id"`
	LanguageID          int32  `json:"language_id"`
	Name                string `json:"name"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

func (q *Queries) CreateContentManagementBannerLanguage(ctx context.Context, arg CreateContentManagementBannerLanguageParams) (int32, error) {
	row := q.db.QueryRow(ctx, createContentManagementBannerLanguage,
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

const deleteContentManagementBannerLanguage = `-- name: DeleteContentManagementBannerLanguage :one
DELETE FROM content_management_banner_languages WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteContentManagementBannerLanguage(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteContentManagementBannerLanguage, id)
	err := row.Scan(&id)
	return id, err
}

const selectAllContentManagementBannerLanguages = `-- name: SelectAllContentManagementBannerLanguages :many
SELECT id, content_management_id, language_id, name, title, content, created_at FROM content_management_banner_languages
`

func (q *Queries) SelectAllContentManagementBannerLanguages(ctx context.Context) ([]ContentManagementBannerLanguage, error) {
	rows, err := q.db.Query(ctx, selectAllContentManagementBannerLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ContentManagementBannerLanguage{}
	for rows.Next() {
		var i ContentManagementBannerLanguage
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

const selectContentManagementBannerLanguages = `-- name: SelectContentManagementBannerLanguages :one
SELECT id, content_management_id, language_id, name, title, content, created_at FROM content_management_banner_languages WHERE  id = $1
`

func (q *Queries) SelectContentManagementBannerLanguages(ctx context.Context, id int32) (ContentManagementBannerLanguage, error) {
	row := q.db.QueryRow(ctx, selectContentManagementBannerLanguages, id)
	var i ContentManagementBannerLanguage
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

const updateContentManagementBannerLanguage = `-- name: UpdateContentManagementBannerLanguage :one
UPDATE content_management_banner_languages
SET name = $3, title = $4, content = $5
WHERE id = $1 AND content_management_id = $2
RETURNING id
`

type UpdateContentManagementBannerLanguageParams struct {
	ID                  int32  `json:"id"`
	ContentManagementID int32  `json:"content_management_id"`
	Name                string `json:"name"`
	Title               string `json:"title"`
	Content             string `json:"content"`
}

func (q *Queries) UpdateContentManagementBannerLanguage(ctx context.Context, arg UpdateContentManagementBannerLanguageParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateContentManagementBannerLanguage,
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
