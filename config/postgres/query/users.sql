-- dipake untuk mengecek username ketika create user baru
-- name: GetUserByUsername :one
SELECT id FROM users
WHERE username = $1;

-- dipake untuk delete user
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
