// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: userShoppingListNames.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserShoppingListName = `-- name: CreateUserShoppingListName :one
INSERT INTO user_shopping_list_names (
    user_id,
    name,
    how_obtain,
    frequency,
    type,
    isnotify
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id
`

type CreateUserShoppingListNameParams struct {
	UserID    int32       `json:"user_id"`
	Name      string      `json:"name"`
	HowObtain pgtype.Text `json:"how_obtain"`
	Frequency pgtype.Bool `json:"frequency"`
	Type      pgtype.Text `json:"type"`
	Isnotify  pgtype.Bool `json:"isnotify"`
}

func (q *Queries) CreateUserShoppingListName(ctx context.Context, arg CreateUserShoppingListNameParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUserShoppingListName,
		arg.UserID,
		arg.Name,
		arg.HowObtain,
		arg.Frequency,
		arg.Type,
		arg.Isnotify,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserShoppingListName = `-- name: DeleteUserShoppingListName :one
DELETE FROM user_shopping_list_names WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUserShoppingListName(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserShoppingListName, id)
	err := row.Scan(&id)
	return id, err
}

const selectUserShoppingListNames = `-- name: SelectUserShoppingListNames :many
SELECT id, user_id, name, how_obtain, frequency, type, isnotify, created_at FROM user_shopping_list_names WHERE user_id = $1
`

func (q *Queries) SelectUserShoppingListNames(ctx context.Context, userID int32) ([]UserShoppingListName, error) {
	rows, err := q.db.Query(ctx, selectUserShoppingListNames, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserShoppingListName{}
	for rows.Next() {
		var i UserShoppingListName
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.HowObtain,
			&i.Frequency,
			&i.Type,
			&i.Isnotify,
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

const updateUserShoppingListName = `-- name: UpdateUserShoppingListName :one
UPDATE user_shopping_list_names
SET
    name = $2,
    how_obtain = $3,
    frequency = $4,
    type = $5,
    isnotify = $6,
    modified = NOW()
WHERE id = $1
RETURNING id
`

type UpdateUserShoppingListNameParams struct {
	ID        int32       `json:"id"`
	Name      string      `json:"name"`
	HowObtain pgtype.Text `json:"how_obtain"`
	Frequency pgtype.Bool `json:"frequency"`
	Type      pgtype.Text `json:"type"`
	Isnotify  pgtype.Bool `json:"isnotify"`
}

func (q *Queries) UpdateUserShoppingListName(ctx context.Context, arg UpdateUserShoppingListNameParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateUserShoppingListName,
		arg.ID,
		arg.Name,
		arg.HowObtain,
		arg.Frequency,
		arg.Type,
		arg.Isnotify,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}