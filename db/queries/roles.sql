-- name: CreateRole :one
INSERT INTO roles (
    name,
    malls_add,
    malls_view,
    malls_edit,
    malls_delete,
    stores_add,
    stores_view,
    stores_edit,
    stores_delete,
    products_add,
    products_view,
    products_edit,
    products_delete,
    products_price,
    users_view,
    users_add,
    users_edit,
    users_delete,
    users_role,
    audits_user,
    audits_product,
    store_locations_view,
    store_locations_add,
    store_locations_edit,
    store_locations_delete,
    offers_add,
    offers_edit,
    offers_delete,
    offers_view,
    reports_view,
    brands_view,
    brands_add,
    brands_edit,
    brands_delete,
    units_view,
    units_add,
    units_edit,
    units_delete,
    product_categories_view,
    product_categories_add,
    product_categories_edit,
    product_categories_delete
) VALUES (
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
    $38,
    $39,
    $40,
    $41,
    $42
) RETURNING id;

-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: UpdateRole :one
UPDATE roles
SET
    name = $2,
    malls_add = $3,
    malls_view = $4,
    malls_edit = $5,
    malls_delete = $6,
    stores_add = $7,
    stores_view = $8,
    stores_edit = $9,
    stores_delete = $10,
    products_add = $11,
    products_view = $12,
    products_edit = $13,
    products_delete = $14,
    products_price = $15,
    users_view = $16,
    users_add = $17,
    users_edit = $18,
    users_delete = $19,
    users_role = $20,
    audits_user = $21,
    audits_product = $22,
    store_locations_view = $23,
    store_locations_add = $24,
    store_locations_edit = $25,
    store_locations_delete = $26,
    offers_add = $27,
    offers_edit = $28,
    offers_delete = $29,
    offers_view = $30,
    reports_view = $31,
    brands_view = $32,
    brands_add = $33,
    brands_edit = $34,
    brands_delete = $35,
    units_view = $36,
    units_add = $37,
    units_edit = $38,
    units_delete = $39,
    product_categories_view = $40,
    product_categories_add = $41,
    product_categories_edit = $42,
    product_categories_delete = $43
WHERE id = $1
RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles WHERE id = $1
RETURNING *;
