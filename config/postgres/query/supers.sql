-- name: GetSuperByUsername :one
SELECT supers.id, user_id, username, password, name FROM users
INNER JOIN supers ON supers.user_id = users.id
WHERE username = $1;
