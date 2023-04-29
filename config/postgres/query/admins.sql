-- name: CreateAdmin :execresult
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($1, $2, $3)
  RETURNING id
)
INSERT INTO admins (user_id)
SELECT id FROM added_user;

-- name: GetAdmins :many
SELECT admins.id, user_id, username, name FROM admins
INNER JOIN users ON admins.user_id = users.id;

-- name: GetAdminByUserId :one
SELECT admins.id, user_id, username, name FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE user_id = $1;

-- name: GetAdminByUsername :one
SELECT admins.id, user_id, username, name FROM users
INNER JOIN admins ON admins.user_id = users.id
WHERE username = $1;

-- name: DeleteAdmin :exec
DELETE FROM admins WHERE user_id = $1;
