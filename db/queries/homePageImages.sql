-- name: CreateHomePageImage :one
INSERT INTO home_page_images (name, path)
VALUES ($1, $2)
RETURNING id;

-- name: SelectHomePageImages :many
SELECT * FROM home_page_images;

-- name: UpdateHomePageImage :one
UPDATE home_page_images
SET
    name = $2,
    path = $3
WHERE id = $1
RETURNING id;

-- name: DeleteHomePageImage :one
DELETE FROM home_page_images WHERE id = $1
RETURNING id;
