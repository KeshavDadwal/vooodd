-- name: CreateUser :one
INSERT INTO users (
    first_name,
    last_name,
    gender,
    dob,
    address,
    city,
    state,
    country_id,
    zip,
    mobile_no,
    email,
    username,
    password,
    user_type,
    latitude,
    longitude,
    photo,
    facebook_id,
    google_id
) VALUES (
    $1,
    $2,
    $3::gender_enum,
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
    $14,
    $15,
    $16,
    $17,
    $18,
    $19
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
    first_name = $2,
    last_name = $3,
    gender = $4::gender_enum,
    dob = $5,
    address = $6,
    city = $7,
    state = $8,
    country_id = $9,
    zip = $10,
    mobile_no = $11,
    email = $12,
    username = $13,
    password = $14,
    user_type = $15,
    latitude = $16,
    longitude = $17,
    photo = $18,
    facebook_id = $19,
    google_id = $20
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1
RETURNING *;
