-- name: CreateHomePageIcon :one
INSERT INTO home_page_icons (name, path, type, value)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectHomePageIcons :one
SELECT * FROM home_page_icons where id = $1;

-- name: SelectAllHomePageIcons :many
SELECT * FROM home_page_icons;

-- name: UpdateHomePageIcon :one
UPDATE home_page_icons
SET
    name = $2,
    path = $3,
    type = $4,
    value = $5
WHERE id = $1
RETURNING id;

-- name: DeleteHomePageIcon :one
DELETE FROM home_page_icons WHERE id = $1
RETURNING id;
