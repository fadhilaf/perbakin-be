-- name: CreateResult :one
INSERT INTO results (shooter_id)
VALUES ($1)
RETURNING id, shooter_id, failed, stage, created_at, updated_at;

-- name: GetResultByShooterId :one
SELECT id, shooter_id, failed, stage, created_at, updated_at
FROM results
WHERE shooter_id = $1;
