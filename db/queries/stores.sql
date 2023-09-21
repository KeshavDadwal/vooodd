-- name: CreateStore :one
INSERT INTO stores (name, store_no, description, logo, phone_no, isdeleted, isblocked, user_id, icon)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: SelectStores :many
SELECT * FROM stores WHERE isdeleted = false AND isblocked = false;

-- name: UpdateStore :one
UPDATE stores
SET name = $2, store_no = $3, description = $4, logo = $5, phone_no = $6, isdeleted = $7, isblocked = $8, user_id = $9, icon = $10
WHERE id = $1
RETURNING id;

-- name: DeleteStore :one
DELETE FROM stores WHERE id = $1
RETURNING id;
