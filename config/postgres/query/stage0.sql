-- name: CreateStage0 :one
INSERT INTO stage0_results (result_id)
VALUES ($1)
RETURNING id, result_id, status, series1, series2, series3, series4, series5, checkmarks, shooter_sign, scorer_sign, created_at, updated_at;

-- name: GetStage0ById :one
SELECT 
  stage0_results.id, 
  stage0_results.result_id, 
  stage0_results.status, 
  stage0_results.series1, 
  stage0_results.series2, 
  stage0_results.series3, 
  stage0_results.series4, 
  stage0_results.series5,
  stage0_results.checkmarks,
  stage0_results.shooter_sign,
  stage0_results.scorer_sign,
  stage0_results.created_at,
  stage0_results.updated_at
FROM stage0_results
WHERE stage0_results.id = $1;

-- name: GetStage0RelationByResultId :one
SELECT 
  stage0_results.id, 
  stage0_results.result_id
FROM stage0_results
WHERE stage0_results.result_id = $1;

-- name: GetStage0Status :one
SELECT 
  stage0_results.status
FROM stage0_results
WHERE stage0_results.id = $1;

-- (scorer role)
-- name: UpdateStage0Checkmarks :one
UPDATE stage0_results
SET checkmarks = $2, updated_at = NOW()
WHERE id = $1 
RETURNING checkmarks;

-- (scorer role)
-- name: UpdateStage0NextSeries :exec
UPDATE stage0_results
SET status = $2, updated_at = NOW()
WHERE id = $1;

-- (scorer role)
-- name: UpdateStage0FinishSuccess :exec
WITH  updated_stage0 AS (
  UPDATE stage0_results
  SET status = '6', shooter_sign = $2, scorer_sign = $3, updated_at = NOW()
  WHERE stage0_results.id = $1
  RETURNING result_id
)
UPDATE results
SET stage = '1', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage0);

-- (scorer role)
-- name: UpdateStage0FinishFailed :exec
WITH updated_stage0 AS (
  UPDATE stage0_results
  SET status = '6', shooter_sign = $2, scorer_sign = $3, updated_at = NOW()
  WHERE stage0_results.id = $1
  RETURNING result_id
)
UPDATE results
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage0);

-- (scorer role) dibuat by id, utk update stage 
-- name: UpdateResultNextStage :exec
UPDATE results 
SET stage = $2, updated_at = NOW()
WHERE id = $1;

-- (scorer role)
-- name: UpdateStage0Series1 :one
UPDATE stage0_results
SET series1 = $2, updated_at = NOW()
WHERE id = $1
RETURNING series1;

-- (scorer role)
-- name: UpdateStage0Series2 :one
UPDATE stage0_results
SET series2 = $2, updated_at = NOW()
WHERE id = $1
RETURNING series2;

-- (scorer role)
-- name: UpdateStage0Series3 :one
UPDATE stage0_results
SET series3 = $2, updated_at = NOW()
WHERE id = $1
RETURNING series3;

-- (scorer role)
-- name: UpdateStage0Series4 :one
UPDATE stage0_results
SET series4 = $2, updated_at = NOW()
WHERE id = $1 
RETURNING series4;

-- (scorer role)
-- name: UpdateStage0Series5 :one
UPDATE stage0_results
SET series5 = $2, updated_at = NOW()
WHERE id = $1 
RETURNING series5;

-- (admin-super role)
-- name: UpdateStage0 :one
UPDATE stage0_results
SET 
  status = $2, 
  series1 = $3, 
  series2 = $4,
  series3 = $5,
  series4 = $6,
  series5 = $7,
  checkmarks = $8,
  updated_at = NOW()
WHERE id = $1 
RETURNING id, result_id, status, series1, series2, series3, series4, series5, checkmarks, created_at, updated_at;

-- (admin-super role)
-- name: UpdateStage0Signs :one
UPDATE stage0_results
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at;

-- (admin-super role)
-- name: DeleteStage0 :exec
WITH deleted_stage0 AS (
  DELETE FROM stage0_results
  WHERE stage0_results.id = $1
  RETURNING result_id
)
UPDATE results 
SET stage = '0', updated_at = NOW()
WHERE id = (SELECT result_id FROM deleted_stage0);
