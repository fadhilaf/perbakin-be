-- name: CreateExam :one
INSERT INTO exams (super_id, name, location, organizer, begin, finish)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, super_id, name, location, organizer, begin, finish;

-- name: GetExamById :one 
SELECT id, super_id, name, location, organizer, begin, finish 
FROM exams 
WHERE id = $1;

-- name: GetExamByName :one
SELECT id, super_id, name, location, organizer, begin, finish 
FROM exams 
WHERE name = $1;

-- name: GetExamBySuperId :many
SELECT id, super_id, name, location, organizer, begin, finish 
FROM exams 
WHERE super_id = $1;

-- name: GetAllExams :many
SELECT id, super_id, name, location, organizer, begin, finish FROM exams;

-- name: UpdateExam :one
UPDATE exams 
SET 
  name = $2, 
  location = $3, 
  organizer = $4, 
  begin = $5, 
  finish = $6 
WHERE id = $1
RETURNING id, super_id, name, location, organizer, begin, finish;

-- name: DeleteExam :exec
DELETE FROM exams WHERE id = $1;
