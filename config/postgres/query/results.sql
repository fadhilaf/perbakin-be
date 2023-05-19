-- (all role)
-- name: CreateResult :one
INSERT INTO results (shooter_id)
VALUES ($1)
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- name: GetResultByShooterId :one
SELECT id, shooter_id, failed, stage, created_at, updated_at
FROM results
WHERE shooter_id = $1;

-- name: GetResultRelationByShooterId :one
SELECT id, shooter_id
FROM results
WHERE shooter_id = $1;

-- (admin-super role) dibuat by shooter id, kareno shooter dan result itu 1:1
-- name: UpdateResultByShooterId :one
UPDATE results 
SET failed = $2, stage = $3, updated_at = NOW()
WHERE shooter_id = $1
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- (admin-super role) dibuat by shooter id, kareno shooter dan result itu 1:1
-- name: DeleteResultByShooterId :exec
DELETE FROM results
WHERE shooter_id = $1;
