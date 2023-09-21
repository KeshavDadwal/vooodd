-- name: CreateMall :one
INSERT INTO malls (
    parent_id,
    name,
    mall_no,
    address,
    city,
    state,
    country_id,
    zip,
    phone_no,
    latitude,
    longitude,
    banner_path,
    isdeleted,
    isblocked,
    timezone,
    monday_start,
    monday_end,
    tuesday_start,
    tuesday_end,
    wednesday_start,
    wednesday_end,
    thursday_start,
    thursday_end,
    friday_start,
    friday_end,
    saturday_start,
    saturday_end,
    sunday_start,
    sunday_end,
    user_id,
    map,
    is_holiday,
    holiday_start_date,
    holiday_end_date,
    holiday_start_time,
    holiday_end_time,
    zone_id,
    ad_link
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
    $14,
    $15,
    $16,
    $17,
    $18,
    $19,
    $20,
    $21,
    $22,
    $23,
    $24,
    $25,
    $26,
    $27,
    $28,
    $29,
    $30,
    $31,
    $32,
    $33,
    $34,
    $35,
    $36,
    $37,
    $38
)
RETURNING id;


-- name: SelectMalls :many
SELECT * FROM malls;

-- name: UpdateMall :one
UPDATE malls
SET
    parent_id = $2,
    name = $3,
    mall_no = $4,
    address = $5,
    city = $6,
    state = $7,
    country_id = $8,
    zip = $9,
    phone_no = $10,
    latitude = $11,
    longitude = $12,
    banner_path = $13,
    timezone = $14,
    monday_start = $15,
    monday_end = $16,
    tuesday_start = $17,
    tuesday_end = $18,
    wednesday_start = $19,
    wednesday_end = $20,
    thursday_start = $21,
    thursday_end = $22,
    friday_start = $23,
    friday_end = $24,
    saturday_start = $25,
    saturday_end = $26,
    sunday_start = $27,
    sunday_end = $28,
    user_id = $29,
    map = $30,
    is_holiday = $31,
    holiday_start_date = $32,
    holiday_end_date = $33,
    holiday_start_time = $34,
    holiday_end_time = $35,
    zone_id = $36,
    ad_link = $37
WHERE id = $1
RETURNING *;


-- name: DeleteMall :one
DELETE FROM malls WHERE id = $1
RETURNING id;

