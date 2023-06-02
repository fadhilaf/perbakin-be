-- untuk membuat exam (super role)
-- name: CreateExam :one
INSERT INTO exams (super_id, name, location, organizer, begin, finish)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, super_id, name, location, organizer, begin, finish, active, created_at, updated_at;

-- untuk mengambil seluruh exam (super role)
-- name: GetAllExams :many
SELECT exams.name, users.name as super, location, organizer, begin, finish, active
FROM exams JOIN supers ON exams.super_id = supers.id 
JOIN users ON supers.user_id = users.id;

-- untuk mengambil seluruh exam (super role)
-- name: GetExamsBySuperId :many
SELECT id, name, location, organizer, begin, finish, active
FROM exams 
WHERE super_id = $1;

-- untuk mengambil exam berdasarkan nama untuk cek nama sudah dipakai blum (super role)
-- name: GetExamByName :one
SELECT id
FROM exams 
WHERE name = $1;

-- untuk mengambil data relasi exam (all role)
-- name: GetExamRelationById :one
SELECT id, super_id 
FROM exams 
WHERE id = $1;

-- untuk mengambil satu data exam (super role)
-- name: GetExamById :one 
SELECT id, super_id, name, location, organizer, begin, finish, active, created_at, updated_at 
FROM exams 
WHERE id = $1;

-- untuk memperbarui exam (super role)
-- name: UpdateExam :one
UPDATE exams 
SET 
  name = $2, 
  location = $3, 
  organizer = $4, 
  begin = $5, 
  finish = $6, 
  updated_at = NOW()
WHERE id = $1
RETURNING id, super_id, name, location, organizer, begin, finish, created_at, updated_at;

-- untuk mengubah status exam (super role)
-- name: UpdateExamStatus :one
UPDATE exams 
  SET active = CASE 
    WHEN active = true 
      THEN false 
      ELSE true 
    END, 
  updated_at = NOW()
WHERE id = $1
RETURNING active;

-- untuk menghapus exam (super role)
-- name: DeleteExam :exec
DELETE FROM exams WHERE id = $1;
