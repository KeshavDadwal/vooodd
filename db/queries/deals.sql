-- name: CreateDeal :one
INSERT INTO deals (parent_id, lft, rght, name)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: SelectDeals :one
SELECT * FROM deals where id = $1;

-- name: SelectAllDeals :many
SELECT * FROM deals;


-- name: UpdateDeal :one
UPDATE deals
SET
    parent_id = $2,
    lft = $3,
    rght = $4,
    name = $5
WHERE id = $1
RETURNING id;

-- name: DeleteDeal :one
DELETE FROM deals WHERE id = $1
RETURNING id;
