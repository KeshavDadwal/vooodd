// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: userMallPermissions.sql

package db

import (
	"context"
)

const createUserMallPermission = `-- name: CreateUserMallPermission :one
INSERT INTO user_mall_permissions (
    user_id,
    mall_id
)
VALUES (
    $1,
    $2
)
RETURNING id
`

type CreateUserMallPermissionParams struct {
	UserID int32 `json:"user_id"`
	MallID int32 `json:"mall_id"`
}

func (q *Queries) CreateUserMallPermission(ctx context.Context, arg CreateUserMallPermissionParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUserMallPermission, arg.UserID, arg.MallID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserMallPermission = `-- name: DeleteUserMallPermission :one
DELETE FROM user_mall_permissions
WHERE
    user_id = $1
    AND mall_id = $2
    RETURNING id
`

type DeleteUserMallPermissionParams struct {
	UserID int32 `json:"user_id"`
	MallID int32 `json:"mall_id"`
}

func (q *Queries) DeleteUserMallPermission(ctx context.Context, arg DeleteUserMallPermissionParams) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserMallPermission, arg.UserID, arg.MallID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserMallPermissionsByUserId = `-- name: GetUserMallPermissionsByUserId :one
SELECT
    id,
    user_id,
    mall_id,
    created_at
FROM
    user_mall_permissions
WHERE
    user_id = $1
`

func (q *Queries) GetUserMallPermissionsByUserId(ctx context.Context, userID int32) (UserMallPermission, error) {
	row := q.db.QueryRow(ctx, getUserMallPermissionsByUserId, userID)
	var i UserMallPermission
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MallID,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserMallPermission = `-- name: UpdateUserMallPermission :one
UPDATE user_mall_permissions
SET
    mall_id = $2
WHERE
    user_id = $1
    RETURNING id
`

type UpdateUserMallPermissionParams struct {
	UserID int32 `json:"user_id"`
	MallID int32 `json:"mall_id"`
}

func (q *Queries) UpdateUserMallPermission(ctx context.Context, arg UpdateUserMallPermissionParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateUserMallPermission, arg.UserID, arg.MallID)
	var id int32
	err := row.Scan(&id)
	return id, err
}