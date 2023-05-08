-- name: GetUserByUsername :one
SELECT id, username, password, name FROM users
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
