-- name: CreateStage0 :one
WITH added_stage0 AS (
  INSERT INTO stage0_results (result_id)
  VALUES ($1)
  RETURNING id, result_id, status, series1, series2, series3, series4, series5, shooter_sign, scorer_sign, created_at, updated_at
), updated_result AS (
  UPDATE results
  SET stage = '0', updated_at = NOW()
  WHERE id = $1
)
SELECT 
  added_stage0.id, 
  added_stage0.result_id, 
  added_stage0.status, 
  added_stage0.series1, 
  added_stage0.series2, 
  added_stage0.series3, 
  added_stage0.series4, 
  added_stage0.series5, 
  added_stage0.shooter_sign,
  added_stage0.scorer_sign,
  added_stage0.created_at, 
  added_stage0.updated_at
FROM added_stage0;

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
-- name: UpdateStage0NextSeries :one
UPDATE stage0_results
SET status = $2, updated_at = NOW()
WHERE id = $1 
RETURNING status;

-- (scorer role)
-- name: UpdateStage0Finish :one
UPDATE stage0_results
SET status = '6', shooter_sign = $2, scorer_sign = $3, updated_at = NOW()
WHERE id = $1
RETURNING status, shooter_sign, scorer_sign;

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
