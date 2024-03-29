-- (all role)
-- name: CreateStage0 :one
INSERT INTO stage0_results (result_id)
VALUES ($1)
RETURNING 
  id, 
  result_id, 
  status, 
  series1, 
  series2, 
  series3, 
  series4, 
  series5, 
  checkmarks, 
  shooter_sign, 
  scorer_sign, 
  created_at, 
  updated_at;

-- (all role)
-- name: GetStage0ById :one
SELECT 
  id, 
  result_id, 
  status, 
  series1, 
  series2, 
  series3, 
  series4, 
  series5,
  checkmarks,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM stage0_results
WHERE id = $1;

-- (all role)
-- name: GetStage0RelationByResultId :one
SELECT 
  id, 
  result_id
FROM stage0_results
WHERE result_id = $1;

-- name: GetStage0Status :one
SELECT 
  status
FROM stage0_results
WHERE id = $1;

-- name: GetStage0Series1 :one
SELECT 
  series1
FROM stage0_results
WHERE id = $1;

-- name: GetStage0Series2 :one
SELECT 
  series2
FROM stage0_results
WHERE id = $1;

-- name: GetStage0Series3 :one
SELECT 
  series3
FROM stage0_results
WHERE id = $1;

-- name: GetStage0Series4 :one
SELECT 
  series4
FROM stage0_results
WHERE id = $1;

-- name: GetStage0Series5 :one
SELECT 
  series5
FROM stage0_results
WHERE id = $1;

-- name: FinishStage0 :exec
WITH updated_stage0 AS (
  UPDATE stage0_results
  SET status = '6', updated_at = NOW()
  WHERE stage0_results.id = $1
  RETURNING result_id
)
UPDATE results 
SET stage = '1', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage0); 

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
  DELETE FROM stage0_results
  WHERE stage0_results.id = $1
  RETURNING result_id;
