-- name: CreateContentManagementMobile :one
INSERT INTO content_management_mobiles (name, title, link, link_text, content, isblocked, isdeleted)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: SelectContentManagementMobiles :one
SELECT * FROM content_management_mobiles WHERE id = $1;

-- name: SelectAllContentManagementMobiles :many
SELECT * FROM content_management_mobiles;

-- name: UpdateContentManagementMobile :one
UPDATE content_management_mobiles
SET name = $2, title = $3, link = $4, link_text = $5, content = $6, isblocked = $7, isdeleted = $8
WHERE id = $1
RETURNING id;

-- name: DeleteContentManagementMobile :one
DELETE FROM content_management_mobiles WHERE id = $1
RETURNING id;
