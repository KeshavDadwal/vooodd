-- name: CreateOffer :one
INSERT INTO offers (deal_no, store_id, product_id, name, path, description, isblocked, isdeleted, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: SelectOffers :many
SELECT * FROM offers;

-- name: UpdateOffer :one
UPDATE offers
SET
    deal_no = $2,
    store_id = $3,
    product_id = $4,
    name = $5,
    path = $6,
    description = $7,
    isblocked = $8,
    isdeleted = $9,
    user_id = $10
WHERE id = $1
RETURNING id;

-- name: DeleteOffer :one
DELETE FROM offers WHERE id = $1
RETURNING id;

