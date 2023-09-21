-- name: CreateUserDevice :one
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
RETURNING id;


-- name: SelectUserDevices :many
SELECT * FROM user_devices WHERE user_id = $1;


-- name: UpdateUserDevice :one
UPDATE user_devices
SET
    device_id = $2,
    device_type = $3::deviceType_enum,
    wuid = $4,
    device_token = $5
WHERE id = $1
RETURNING id;

-- name: DeleteUserDevice :one
DELETE FROM user_devices WHERE id = $1
RETURNING id;

