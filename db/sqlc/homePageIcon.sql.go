// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: homePageIcon.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createHomePageIcon = `-- name: CreateHomePageIcon :one
INSERT INTO home_page_icons (name, path, type, value)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateHomePageIconParams struct {
	Name  string      `json:"name"`
	Path  string      `json:"path"`
	Type  pgtype.Text `json:"type"`
	Value string      `json:"value"`
}

func (q *Queries) CreateHomePageIcon(ctx context.Context, arg CreateHomePageIconParams) (int32, error) {
	row := q.db.QueryRow(ctx, createHomePageIcon,
		arg.Name,
		arg.Path,
		arg.Type,
		arg.Value,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteHomePageIcon = `-- name: DeleteHomePageIcon :one
DELETE FROM home_page_icons WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteHomePageIcon(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteHomePageIcon, id)
	err := row.Scan(&id)
	return id, err
}

const selectAllHomePageIcons = `-- name: SelectAllHomePageIcons :many
SELECT id, name, path, type, value, isdeleted, isblocked, created_at FROM home_page_icons
`

func (q *Queries) SelectAllHomePageIcons(ctx context.Context) ([]HomePageIcon, error) {
	rows, err := q.db.Query(ctx, selectAllHomePageIcons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []HomePageIcon{}
	for rows.Next() {
		var i HomePageIcon
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Path,
			&i.Type,
			&i.Value,
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

const selectHomePageIcons = `-- name: SelectHomePageIcons :one
SELECT id, name, path, type, value, isdeleted, isblocked, created_at FROM home_page_icons where id = $1
`

func (q *Queries) SelectHomePageIcons(ctx context.Context, id int32) (HomePageIcon, error) {
	row := q.db.QueryRow(ctx, selectHomePageIcons, id)
	var i HomePageIcon
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Path,
		&i.Type,
		&i.Value,
		&i.Isdeleted,
		&i.Isblocked,
		&i.CreatedAt,
	)
	return i, err
}

const updateHomePageIcon = `-- name: UpdateHomePageIcon :one
UPDATE home_page_icons
SET
    name = $2,
    path = $3,
    type = $4,
    value = $5
WHERE id = $1
RETURNING id
`

type UpdateHomePageIconParams struct {
	ID    int32       `json:"id"`
	Name  string      `json:"name"`
	Path  string      `json:"path"`
	Type  pgtype.Text `json:"type"`
	Value string      `json:"value"`
}

func (q *Queries) UpdateHomePageIcon(ctx context.Context, arg UpdateHomePageIconParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateHomePageIcon,
		arg.ID,
		arg.Name,
		arg.Path,
		arg.Type,
		arg.Value,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
