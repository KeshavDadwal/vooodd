-- name: CreateContentManagementBanner :one
INSERT INTO content_management_banners (content_management_id, banner, banner_text, banner_link)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectContentManagementBanners :one
SELECT * FROM content_management_banners WHERE id = $1;

-- name: SelectAllContentManagementBanners :many
SELECT * FROM content_management_banners;

-- name: UpdateContentManagementBanner :one
UPDATE content_management_banners
SET banner = $2, banner_text = $3, banner_link = $4
WHERE id = $1
RETURNING id;

-- name: DeleteContentManagementBanner :one
DELETE FROM content_management_banners WHERE id = $1
RETURNING id;
