-- name: CreateHomeBanner :one
INSERT INTO home_banners (name, url_type, description, param, image, isdefault, sequence, start_date, end_date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: SelectHomeBanners :one
SELECT * FROM home_banners where id = $1;

-- name: SelectAllHomeBanners :many
SELECT * FROM home_banners;


-- name: UpdateHomeBanner :one
UPDATE home_banners
SET
    name = $2,
    url_type = $3,
    description = $4,
    param = $5,
    image = $6,
    isdefault = $7,
    sequence = $8,
    start_date = $9,
    end_date = $10
WHERE id = $1
RETURNING id;

-- name: DeleteHomeBanner :one
DELETE FROM home_banners WHERE id = $1
RETURNING id;
