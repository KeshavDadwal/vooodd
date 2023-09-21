-- name: CreateUserRole :one
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, $2)
RETURNING id;

-- name: SelectUserRoles :many
SELECT * FROM user_roles WHERE user_id = $1;

-- name: UpdateUserRole :one
UPDATE user_roles
SET user_id = $2, role_id = $3
WHERE id = $1
RETURNING id;

-- name: DeleteUserRole :one
DELETE FROM user_roles WHERE id = $1
RETURNING id;

