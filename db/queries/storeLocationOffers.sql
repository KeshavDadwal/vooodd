-- name: CreateStoreLocationOffer :one
INSERT INTO store_location_offers (
    store_location_id,
    product_id,
    offer_id,
    price,
    offered_price,
    is_percent,
    start_date,
    end_date,
    sequence,
    isblocked,
    isdeleted,
    priority,
    weight,
    unit_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14
)
RETURNING id;

-- name: SelectStoreLocationOffers :many
SELECT * FROM store_location_offers;

-- name: UpdateStoreLocationOffer :one
UPDATE store_location_offers
SET
    store_location_id = $2,
    product_id = $3,
    offer_id = $4,
    price = $5,
    offered_price = $6,
    is_percent = $7,
    start_date = $8,
    end_date = $9,
    sequence = $10,
    isblocked = $11,
    isdeleted = $12,
    priority = $13,
    weight = $14,
    unit_id = $15
WHERE id = $1
RETURNING *;

-- name: DeleteStoreLocationOffer :one
DELETE FROM store_location_offers WHERE id = $1
RETURNING id;


