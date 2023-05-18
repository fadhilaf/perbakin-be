-- name: CreateStage0 :one
WITH added_stage0 AS (
  INSERT INTO stage0_results (result_id)
  VALUES ($1)
  RETURNING id, result_id, status
), updated_result AS (
  UPDATE results
  SET stage = '0', updated_at = NOW()
  WHERE id = $1
  RETURNING id, shooter_id, failed, stage, created_at, updated_at
)
SELECT 
  updated_result.id, 
  updated_result.shooter_id, 
  updated_result.failed, 
  updated_result.stage, 
  updated_result.created_at, 
  updated_result.updated_at, 
  added_stage0.id, 
  added_stage0.result_id, 
  added_stage0.status
FROM updated_result
INNER JOIN added_stage0 ON added_stage0.result_id = updated_result.id;

-- name: GetStage0 :one
SELECT 
  stage0_results.id, 
  stage0_results.result_id, 
  stage0_results.status, 
  stage0_results.series1, 
  stage0_results.series2, 
  stage0_results.series3, 
  stage0_results.series4, 
  stage0_results.series5
FROM stage0_results
WHERE stage0_results.result_id = $1;

-- name: UpdateStage0Series1 :one
UPDATE stage0_results
SET series1 = $2, updated_at = NOW()
WHERE result_id = $1
RETURNING series1;

-- name: UpdateStage0NextSeries2 :one
UPDATE stage0_results
SET series1 = $2, status = 2, updated_at = NOW()
WHERE result_id = $1 
RETURNING series1, status;

-- name: UpdateStage0Series2 :one
UPDATE stage0_results
SET series2 = $2, updated_at = NOW()
WHERE result_id = $1
RETURNING series2;

-- name: UpdateStage0NextSeries3 :one
UPDATE stage0_results
SET series2 = $2, status = 3, updated_at = NOW()
WHERE result_id = $1 
RETURNING series2, status;

-- name: UpdateStage0Series3 :one
UPDATE stage0_results
SET series3 = $2, updated_at = NOW()
WHERE result_id = $1
RETURNING series3;

-- name: UpdateStage0NextSeries4 :one
UPDATE stage0_results
SET series3 = $2, status = 4, updated_at = NOW()
WHERE result_id = $1 
RETURNING series3, status;

-- name: UpdateStage0Series4 :one
UPDATE stage0_results
SET series4 = $2, updated_at = NOW()
WHERE result_id = $1 
RETURNING series4;

-- name: UpdateStage0NextSeries5 :one
UPDATE stage0_results
SET series4 = $2, status = 5, updated_at = NOW()
WHERE result_id = $1 
RETURNING series4, status;

-- name: UpdateStage0Series5 :one
UPDATE stage0_results
SET series5 = $2, updated_at = NOW()
WHERE result_id = $1 
RETURNING series5;

-- name: UpdateStage0NextSeries6 :one
UPDATE stage0_results
SET series5 = $2, status = 6, updated_at = NOW()
WHERE result_id = $1 
RETURNING series5, status;
