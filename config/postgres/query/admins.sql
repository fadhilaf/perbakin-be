-- untuk ngebuat admin (super role) TODO: return sebanyak get admin by id
-- name: CreateAdmin :one
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($2, $3, $4)
  RETURNING id
), added_admin AS (
  INSERT INTO admins (user_id, exam_id)
  SELECT id, $1 FROM added_user
  RETURNING id
)
SELECT admins.id, user_id, exam_id, username, name, created_at, updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE admins.id = (
  SELECT id FROM added_admin
);

-- untuk ngambil data display seluruh admin (all role)
-- name: GetAllAdmins :many
SELECT admins.id, exams.name AS exam, users.name AS name, users.created_at, users.updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id
INNER JOIN exams ON admins.exam_id = exams.id;

-- untuk ngambil data akun seluruh admin dalam satu exam (super role)
-- name: GetAdminsByExamId :many
SELECT admins.id, user_id, exam_id, username, name, created_at, updated_at FROM admins 
INNER JOIN users ON admins.user_id = users.id
WHERE exam_id = $1;

-- untuk ngambil data relasi admin berdasarkan user id (all role)
-- name: GetAdminRelationByUserId :one
SELECT admins.id, user_id, exam_id FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE user_id = $1;

-- untuk ngambil data display admin berdasarkan username (admin role)
-- name: GetAdminByUsername :one
SELECT admins.id, user_id, exam_id, username, password, name, created_at, updated_at FROM users
INNER JOIN admins ON admins.user_id = users.id
WHERE username = $1;

-- untuk ngambil data akun admin berdasarkan id (super role)
-- name: GetAdminById :one
SELECT admins.id, user_id, exam_id, username, name, created_at, updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE admins.id = $1;

-- untuk update data akun admin (super role)
-- name: UpdateAdmin :one
WITH updated_user AS (
  UPDATE users 
  SET username = $2, password = $3, name = $4, updated_at = NOW()
  WHERE users.id = (
    SELECT user_id FROM admins 
    WHERE admins.id = $1
  )
  RETURNING id
)
SELECT admins.id, user_id, exam_id, username, name, created_at, updated_at FROM admins
INNER JOIN users ON admins.user_id = users.id
WHERE user_id = (
  SELECT id FROM updated_user
);

-- low prio
-- name: UpdateAdminName :one
UPDATE users 
SET name = $2, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM admins 
  WHERE admins.id = $1
)
RETURNING id;

-- low prio
-- name: UpdateAdminPassword :one
UPDATE users 
SET password = $2, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM admins 
  WHERE admins.id = $1
)
RETURNING id;
