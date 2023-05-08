-- name: CreateAdmin :execresult
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($1, $2, $3)
  RETURNING id
)
INSERT INTO admins (user_id)
SELECT id FROM added_user;

-- name: GetAllAdmins :many
SELECT admins.id, name, created_at, updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id;

-- name: GetAdminData :one
SELECT admins.id, name, created_at, updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE admins.id = $1;

-- name: GetAdminByUserId :one
SELECT admins.id, user_id, username, name FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE user_id = $1;

-- name: GetAdminByUsername :one
SELECT admins.id, user_id, username, password, name FROM users
INNER JOIN admins ON admins.user_id = users.id
WHERE username = $1;

-- name: GetAdmin :one
SELECT admins.id, user_id, username, name, created_at, updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE admins.id = $1;

-- name: UpdateAdmin :exec
UPDATE users 
SET username = $2, password = $3, name = $4, updated_at = NOW()
WHERE users.id = (
  SELECT user_id FROM admins 
  WHERE admins.id = $1
);

-- name: UpdateAdminName :exec
UPDATE users 
SET name = $2, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM admins 
  WHERE admins.id = $1
);

-- name: UpdateAdminPassword :exec
UPDATE users 
SET password = $2, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM admins 
  WHERE admins.id = $1
);
