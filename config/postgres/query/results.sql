-- (all role)
-- name: CreateResult :one
INSERT INTO results (shooter_id)
VALUES ($1)
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- (admin-super role)
-- name: GetResultsByExamId :many
SELECT shooters.id, shooters.name, shooters.province, shooters.club, results.failed, results.stage
FROM results 
JOIN shooters ON results.shooter_id = shooters.id 
JOIN exams ON shooters.exam_id = exams.id
WHERE exams.id = $1;

-- (all role)
-- name: GetResultsByScorerId :many
SELECT shooters.id, shooters.name, shooters.province, shooters.club, results.failed, results.stage 
FROM results 
JOIN shooters ON results.shooter_id = shooters.id 
WHERE shooters.scorer_id = $1; 

-- name: GetResultById :one
SELECT id, shooter_id, failed, stage, created_at, updated_at
FROM results
WHERE id = $1;

-- name: GetResultRelationAndStatusByShooterId :one
SELECT id, shooter_id, stage, failed
FROM results
WHERE shooter_id = $1;

-- (admin-super role) dibuat by id
-- name: UpdateResult :one
UPDATE results 
SET failed = $2, stage = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- (admin-super role) utk edge case kalo delete stage yang terakir, mundurin ke stage sebelum
-- name: UpdateResultStage :exec
UPDATE results 
SET stage = $2, updated_at = NOW()
WHERE id = $1;

-- (admin-super role) dibuat by id
-- name: DeleteResult :exec
DELETE FROM results
WHERE id = $1;
