-- untuk ngebuat scorer (admin-super role) TODO: return sebanyak get scorer by id
-- name: CreateScorer :one
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($2, $3, $4)
  RETURNING id, username, password, name, created_at, updated_at
), added_scorer AS (
  INSERT INTO scorers (user_id, exam_id, image_path)
  SELECT id, $1, COALESCE(sqlc.narg('image_path')::text,'default.png') FROM added_user
  RETURNING id, user_id, exam_id, image_path
)
SELECT added_scorer.id, user_id, exam_id, username, name, image_path, created_at, updated_at 
FROM added_user, added_scorer;

-- untuk ngambil data display seluruh scorer (all role)
-- name: GetAllScorers :many
SELECT exams.name AS exam, users.name AS name, image_path FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
INNER JOIN exams ON scorers.exam_id = exams.id;

-- untuk ngambil data akun seluruh scorer dalam satu exam (admin-super role)
-- name: GetScorersByExamId :many
SELECT scorers.id, name, image_path FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
WHERE exam_id = $1;

-- untuk ngambil data relasi scorer berdasarkan user id (all role)
-- name: GetScorerRelationByUserId :one
SELECT scorers.id, user_id, exam_id, active FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
INNER JOIN exams ON scorers.exam_id = exams.id
WHERE user_id = $1;

-- untuk ngambil data relasi scorer berdasarkan id (all role)
-- name: GetScorerRelationById :one
SELECT scorers.id, user_id, exam_id FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.id = $1;

-- untuk ngambil data display scorer berdasarkan username (scorer role)
-- name: GetScorerByUsername :one
SELECT scorers.id, user_id, exam_id, username, password, users.name, image_path, active, users.created_at, users.updated_at FROM users
INNER JOIN scorers ON scorers.user_id = users.id
INNER JOIN exams ON scorers.exam_id = exams.id
WHERE username = $1;

-- untuk ngambil data akun scorer berdasarkan id (admin-super role)
-- name: GetScorerById :one
SELECT scorers.id, user_id, exam_id, username, name, image_path, created_at, updated_at FROM scorers
INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.id = $1;

-- untuk update data akun admin (super role) TODO: return sebanyak get admin by id
-- name: UpdateScorer :one
WITH updated_user AS (
  UPDATE users 
  SET username = $2, password = COALESCE(sqlc.narg('password')::text,password), name = $3, updated_at = NOW() 
  WHERE users.id = (
    SELECT user_id FROM scorers 
    WHERE scorers.id = $1
  )
  RETURNING id, username, name, created_at, updated_at
)
SELECT scorers.id, user_id, exam_id, username, name, image_path, created_at, updated_at FROM scorers
INNER JOIN updated_user ON scorers.user_id = updated_user.id
WHERE scorers.id = $1;

-- untuk mengupdate foto scorer berdasarkan id (all role)
-- name: UpdateScorerImage :one
WITH updated_user AS (
  UPDATE users 
  SET updated_at = NOW() 
  WHERE users.id = (
    SELECT user_id FROM scorers 
    WHERE scorers.id = $1
  )
)
UPDATE scorers
SET image_path = $2
WHERE scorers.id = $1
RETURNING image_path;
