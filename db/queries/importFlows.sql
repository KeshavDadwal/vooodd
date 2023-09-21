-- name: CreateImportFlow :one
INSERT INTO import_flows (
    file_name,
    no_rows,
    note,
    status
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: SelectAllImportFlowByID :many
SELECT * FROM import_flows;

-- name: UpdateImportFlow :one
UPDATE import_flows
SET
    file_name = $2,
    no_rows = $3,
    note = $4,
    status = $5
WHERE id = $1
RETURNING *;

-- name: DeleteImportFlow :one
DELETE FROM import_flows WHERE id = $1
RETURNING id;
