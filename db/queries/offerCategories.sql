-- name: CreateOfferCategory :one
INSERT INTO offers_categories (deal_id, offer_id)
VALUES ($1, $2)
RETURNING id;

-- name: SelectOfferCategories :many
SELECT * FROM offers_categories;

-- name: UpdateOfferCategory :one
UPDATE offers_categories
SET
    deal_id = $2,
    offer_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteOfferCategory :one
DELETE FROM offers_categories WHERE id = $1
RETURNING id;
