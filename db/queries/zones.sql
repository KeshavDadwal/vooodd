-- name: CreateZone :one
INSERT INTO zones (
    name,
    state
)
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: SelectZones :many
SELECT * FROM zones;

-- name: UpdateZone :one
UPDATE zones
SET
    name = $2,
    state = $3
WHERE id = $1
RETURNING *;

-- name: DeleteZone :one
DELETE FROM zones WHERE id = $1
RETURNING id;

