// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: homePageImages.sql

package db

import (
	"context"
)

const createHomePageImage = `-- name: CreateHomePageImage :one
INSERT INTO home_page_images (name, path)
VALUES ($1, $2)
RETURNING id
`

type CreateHomePageImageParams struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (q *Queries) CreateHomePageImage(ctx context.Context, arg CreateHomePageImageParams) (int32, error) {
	row := q.db.QueryRow(ctx, createHomePageImage, arg.Name, arg.Path)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteHomePageImage = `-- name: DeleteHomePageImage :one
DELETE FROM home_page_images WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteHomePageImage(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteHomePageImage, id)
	err := row.Scan(&id)
	return id, err
}

const selectHomePageImages = `-- name: SelectHomePageImages :many
SELECT id, name, path, isdeleted, isblocked, created_at FROM home_page_images
`

func (q *Queries) SelectHomePageImages(ctx context.Context) ([]HomePageImage, error) {
	rows, err := q.db.Query(ctx, selectHomePageImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []HomePageImage{}
	for rows.Next() {
		var i HomePageImage
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Path,
			&i.Isdeleted,
			&i.Isblocked,
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

const updateHomePageImage = `-- name: UpdateHomePageImage :one
UPDATE home_page_images
SET
    name = $2,
    path = $3
WHERE id = $1
RETURNING id
`

type UpdateHomePageImageParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (q *Queries) UpdateHomePageImage(ctx context.Context, arg UpdateHomePageImageParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateHomePageImage, arg.ID, arg.Name, arg.Path)
	var id int32
	err := row.Scan(&id)
	return id, err
}