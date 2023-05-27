-- (all role)
-- name: CreateResult :one
INSERT INTO results (shooter_id)
VALUES ($1)
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- name: GetResultById :one
SELECT id, shooter_id, failed, stage, created_at, updated_at
FROM results
WHERE id = $1;

-- name: GetResultStageById :one
SELECT stage
FROM results
WHERE id = $1;

-- name: GetResultRelationByShooterId :one
SELECT id, shooter_id
FROM results
WHERE shooter_id = $1;

-- (admin-super role) dibuat by id
-- name: UpdateResult :one
UPDATE results 
SET failed = $2, stage = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- (admin-super role) dibuat by id
-- name: DeleteResult :exec
DELETE FROM results
WHERE id = $1;
