// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: userDevices.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserDevice = `-- name: CreateUserDevice :one
INSERT INTO user_devices (
    user_id,
    device_id,
    device_type,
    wuid,
    device_token
)
VALUES (
    $1,
    $2,
    $3::deviceType_enum,
    $4,
    $5
)
RETURNING id
`

type CreateUserDeviceParams struct {
	UserID      int32          `json:"user_id"`
	DeviceID    string         `json:"device_id"`
	Column3     DevicetypeEnum `json:"column_3"`
	Wuid        pgtype.Text    `json:"wuid"`
	DeviceToken pgtype.Text    `json:"device_token"`
}

func (q *Queries) CreateUserDevice(ctx context.Context, arg CreateUserDeviceParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUserDevice,
		arg.UserID,
		arg.DeviceID,
		arg.Column3,
		arg.Wuid,
		arg.DeviceToken,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserDevice = `-- name: DeleteUserDevice :one
DELETE FROM user_devices WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUserDevice(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserDevice, id)
	err := row.Scan(&id)
	return id, err
}

const selectUserDevices = `-- name: SelectUserDevices :many
SELECT id, user_id, device_id, device_type, wuid, device_token, created_at FROM user_devices WHERE user_id = $1
`

func (q *Queries) SelectUserDevices(ctx context.Context, userID int32) ([]UserDevice, error) {
	rows, err := q.db.Query(ctx, selectUserDevices, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserDevice{}
	for rows.Next() {
		var i UserDevice
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DeviceID,
			&i.DeviceType,
			&i.Wuid,
			&i.DeviceToken,
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

const updateUserDevice = `-- name: UpdateUserDevice :one
UPDATE user_devices
SET
    device_id = $2,
    device_type = $3::deviceType_enum,
    wuid = $4,
    device_token = $5
WHERE id = $1
RETURNING id
`

type UpdateUserDeviceParams struct {
	ID          int32          `json:"id"`
	DeviceID    string         `json:"device_id"`
	Column3     DevicetypeEnum `json:"column_3"`
	Wuid        pgtype.Text    `json:"wuid"`
	DeviceToken pgtype.Text    `json:"device_token"`
}

func (q *Queries) UpdateUserDevice(ctx context.Context, arg UpdateUserDeviceParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateUserDevice,
		arg.ID,
		arg.DeviceID,
		arg.Column3,
		arg.Wuid,
		arg.DeviceToken,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
