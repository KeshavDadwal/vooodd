-- name: CreatePushNotification :one
INSERT INTO push_notifications (
    device_id,
    store_location_id,
    offer_id,
    is_sent,
    is_clicked
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id;

-- name: GetStoreLocationPushNotifications :many
SELECT
    id,
    device_id,
    store_location_id,
    offer_id,
    is_sent,
    is_clicked,
    created_at
FROM
    push_notifications
WHERE
    store_location_id = $1;

-- name: UpdatePushNotificationStatus :one
UPDATE push_notifications
SET
    is_sent = $2,
    is_clicked = $3
WHERE
    id = $1
    RETURNING id;

-- name: DeletePushNotification :one
DELETE FROM push_notifications
WHERE
    id = $1
    RETURNING id;
