// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: languages.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLanguage = `-- name: CreateLanguage :one
INSERT INTO languages (
    name,
    code,
    isdefault,
    isdeleted,
    isblocked
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
        )
RETURNING id, name, code, isdefault, isdeleted, isblocked, created_at
`

type CreateLanguageParams struct {
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Isdefault pgtype.Bool `json:"isdefault"`
	Isdeleted pgtype.Bool `json:"isdeleted"`
	Isblocked pgtype.Bool `json:"isblocked"`
}

func (q *Queries) CreateLanguage(ctx context.Context, arg CreateLanguageParams) (Language, error) {
	row := q.db.QueryRow(ctx, createLanguage,
		arg.Name,
		arg.Code,
		arg.Isdefault,
		arg.Isdeleted,
		arg.Isblocked,
	)
	var i Language
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Isdefault,
		&i.Isdeleted,
		&i.Isblocked,
		&i.CreatedAt,
	)
	return i, err
}

const deleteLanguage = `-- name: DeleteLanguage :one
DELETE FROM languages WHERE id = $1
RETURNING id, name, code, isdefault, isdeleted, isblocked, created_at
`

func (q *Queries) DeleteLanguage(ctx context.Context, id int32) (Language, error) {
	row := q.db.QueryRow(ctx, deleteLanguage, id)
	var i Language
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Isdefault,
		&i.Isdeleted,
		&i.Isblocked,
		&i.CreatedAt,
	)
	return i, err
}

const getAllLanguage = `-- name: GetAllLanguage :many
SELECT id, name, code, isdefault, isdeleted, isblocked, created_at FROM languages
`

func (q *Queries) GetAllLanguage(ctx context.Context) ([]Language, error) {
	rows, err := q.db.Query(ctx, getAllLanguage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Language{}
	for rows.Next() {
		var i Language
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Isdefault,
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

const updateLanguage = `-- name: UpdateLanguage :one
UPDATE languages
SET name = $1, code = $2, isdefault = $3, isdeleted = $4, isblocked = $5
WHERE id = $6
RETURNING id, name, code, isdefault, isdeleted, isblocked, created_at
`

type UpdateLanguageParams struct {
	Name      string      `json:"name"`
	Code      string      `json:"code"`
	Isdefault pgtype.Bool `json:"isdefault"`
	Isdeleted pgtype.Bool `json:"isdeleted"`
	Isblocked pgtype.Bool `json:"isblocked"`
	ID        int32       `json:"id"`
}

func (q *Queries) UpdateLanguage(ctx context.Context, arg UpdateLanguageParams) (Language, error) {
	row := q.db.QueryRow(ctx, updateLanguage,
		arg.Name,
		arg.Code,
		arg.Isdefault,
		arg.Isdeleted,
		arg.Isblocked,
		arg.ID,
	)
	var i Language
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Isdefault,
		&i.Isdeleted,
		&i.Isblocked,
		&i.CreatedAt,
	)
	return i, err
}
