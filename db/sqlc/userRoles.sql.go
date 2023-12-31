// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: userRoles.sql

package db

import (
	"context"
)

const createUserRole = `-- name: CreateUserRole :one
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2)
RETURNING id
`

type CreateUserRoleParams struct {
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}

func (q *Queries) CreateUserRole(ctx context.Context, arg CreateUserRoleParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUserRole, arg.UserID, arg.RoleID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserRole = `-- name: DeleteUserRole :one
DELETE FROM user_roles WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUserRole(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserRole, id)
	err := row.Scan(&id)
	return id, err
}

const selectUserRoles = `-- name: SelectUserRoles :many
SELECT id, user_id, role_id, created_at FROM user_roles WHERE user_id = $1
`

func (q *Queries) SelectUserRoles(ctx context.Context, userID int32) ([]UserRole, error) {
	rows, err := q.db.Query(ctx, selectUserRoles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserRole{}
	for rows.Next() {
		var i UserRole
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.RoleID,
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

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE user_roles
SET user_id = $2, role_id = $3
WHERE id = $1
RETURNING id
`

type UpdateUserRoleParams struct {
	ID     int32 `json:"id"`
	UserID int32 `json:"user_id"`
	RoleID int32 `json:"role_id"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateUserRole, arg.ID, arg.UserID, arg.RoleID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
