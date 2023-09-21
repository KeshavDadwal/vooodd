// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: userEvents.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserEvent = `-- name: CreateUserEvent :one
INSERT INTO user_events (
    user_id,
    name,
    date
)
VALUES (
    $1,
    $2,
    $3
)
RETURNING id
`

type CreateUserEventParams struct {
	UserID int32       `json:"user_id"`
	Name   string      `json:"name"`
	Date   pgtype.Date `json:"date"`
}

func (q *Queries) CreateUserEvent(ctx context.Context, arg CreateUserEventParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUserEvent, arg.UserID, arg.Name, arg.Date)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserEvent = `-- name: DeleteUserEvent :one
DELETE FROM user_events WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUserEvent(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserEvent, id)
	err := row.Scan(&id)
	return id, err
}

const selectUserEvents = `-- name: SelectUserEvents :many
SELECT id, user_id, name, date, created_at FROM user_events WHERE user_id = $1
`

func (q *Queries) SelectUserEvents(ctx context.Context, userID int32) ([]UserEvent, error) {
	rows, err := q.db.Query(ctx, selectUserEvents, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserEvent{}
	for rows.Next() {
		var i UserEvent
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Date,
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

const updateUserEvent = `-- name: UpdateUserEvent :one
UPDATE user_events
SET
    name = $2,
    date = $3
WHERE id = $1
RETURNING id
`

type UpdateUserEventParams struct {
	ID   int32       `json:"id"`
	Name string      `json:"name"`
	Date pgtype.Date `json:"date"`
}

func (q *Queries) UpdateUserEvent(ctx context.Context, arg UpdateUserEventParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateUserEvent, arg.ID, arg.Name, arg.Date)
	var id int32
	err := row.Scan(&id)
	return id, err
}