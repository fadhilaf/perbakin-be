-- untuk ngambil data display seluruh super admin (all role)
-- name: GetAllSupers :many
SELECT supers.id, name, created_at, updated_at FROM supers
INNER JOIN users ON supers.user_id = users.id;

-- untuk ngambil data relasi super admin berdasarkan user id (all role)
-- name: GetSuperRelationByUserId :one
SELECT supers.id, user_id FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE user_id = $1;

-- untuk ngambil data display super admin berdasarkan username (super role)
-- name: GetSuperByUsername :one
SELECT supers.id, user_id, username, password, name, created_at, updated_at FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE username = $1;
