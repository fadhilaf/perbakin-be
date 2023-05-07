-- name: GetSupers :many
SELECT supers.id, user_id, username, password, name FROM supers
INNER JOIN users ON supers.user_id = users.id;

-- name: GetSuperByUserId :one
SELECT supers.id, user_id, username, password, name FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE user_id = $1;

-- name: GetSuperByUsername :one
SELECT supers.id, user_id, username, password, name FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE username = $1;

-- name: UpdateSuper :exec
UPDATE users SET username = $2, name = $3 WHERE id = $1;
