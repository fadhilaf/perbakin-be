-- membuat shooter baru (admin-super role)
-- name: CreateShooter :one
INSERT INTO shooters (scorer_id, name, image_path, province, club)
VALUES ($1, $2, COALESCE(sqlc.narg('image_path')::text,'default.png'), $3, $4)
RETURNING id, scorer_id, name, image_path, province, club, created_at, updated_at;

-- untuk mengambil seluruh shooter (admin-super role)
-- name: GetAllShooters :many
SELECT exams.name AS exam, shooters.name AS name, shooters.image_path, province, club
FROM shooters INNER JOIN scorers ON shooters.scorer_id = scorers.id INNER JOIN exams ON scorers.exam_id = exams.id;

-- untuk mengambil shooter berdasarkan exam_id (admin-super role)
-- name: GetShooterByExamId :many
SELECT shooters.id, scorer_id, users.name AS scorer, shooters.name AS name, shooters.image_path, province, club
FROM shooters INNER JOIN scorers ON shooters.scorer_id = scorers.id INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.exam_id = $1;

-- untuk mengambil shooter berdasarkan scorer_id (all role)
-- name: GetShootersByScorerId :many
SELECT id, name, image_path, province, club
FROM shooters
WHERE scorer_id = $1;

-- untuk mengambil shooter berdasarkan id (admin-super role)
-- name: GetShooterById :one
SELECT id, scorer_id, name, image_path, province, club, created_at, updated_at
FROM shooters
WHERE id = $1;

-- untuk mengambil relasi shooter berdasarkan id (all role)
-- name: GetShooterRelationById :one
SELECT id, scorer_id
FROM shooters
WHERE id = $1;

-- untuk mengupdate shooter berdasarkan id (admin-super role)
-- name: UpdateShooter :one
UPDATE shooters 
SET scorer_id = $2, name = $3, province = $4, club = $5, updated_at = NOW()
WHERE id = $1
RETURNING id, scorer_id, name, province, club, created_at, updated_at;

-- untuk mengupdate foto shooter berdasarkan id (admin-super role)
-- name: UpdateShooterImage :one
UPDATE shooters 
SET image_path = $2, updated_at = NOW()
WHERE id = $1
RETURNING id, scorer_id, name, image_path, province, club, created_at, updated_at;

-- untuk menghapus shooter berdasarkan id (admin-super role)
-- name: DeleteShooter :exec
DELETE FROM shooters
WHERE id = $1;
