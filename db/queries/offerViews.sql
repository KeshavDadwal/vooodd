-- name: CreateOfferView :one
INSERT INTO offer_views (
    product_id,
    offer_id,
    mall_id,
    store_id,
    store_location_id,
    user_id,
    viewed,
    view_date,
    is_clicked
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
    $9
)
RETURNING id;

-- name: SelectOfferViewsByOfferID :many
SELECT * FROM offer_views WHERE offer_id = $1;



-- name: DeleteOfferView :one
DELETE FROM offer_views WHERE id = $1
RETURNING id;

