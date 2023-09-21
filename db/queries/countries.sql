-- name: CreateCountry :one
INSERT INTO countries (
    iso2,
    short_name,
    long_name,
    iso3,
    numcode,
    un_member,
    calling_code,
    cctld
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING id;

-- name: GetCountry :one
SELECT * FROM countries WHERE id = $1;

-- name: GetAllCountry :one
SELECT * FROM countries;


-- name: UpdateCountry :one
UPDATE countries
SET
    iso2 = $2,
    short_name = $3,
    long_name = $4,
    iso3 = $5,
    numcode = $6,
    un_member = $7,
    calling_code = $8,
    cctld = $9
WHERE id = $1
RETURNING *;

-- name: DeleteCountry :one
DELETE FROM countries WHERE id = $1
RETURNING *;
