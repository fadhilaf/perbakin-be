-- name: CreateStage1 :one
INSERT INTO stage1_results (result_id)
VALUES ($1)
RETURNING 
  id, 
  result_id, 
  ( try1 ).status,
  ( try1 ).no1.scores,
  ( try1 ).no1.duration,
  ( try1 ).no2.scores,
  ( try1 ).no2.duration,
  ( try1 ).no3.scores,
  ( try1 ).no3.duration,
  ( try1 ).no4.scores,
  ( try1 ).no4.duration,
  ( try1 ).no5.scores,
  ( try1 ).no5.duration,
  ( try1 ).no6.scores,
  ( try1 ).no6.duration,
  ( try1 ).checkmarks,
  created_at,
  updated_at;

-- (all role)
-- name: GetStage1try1ById :one
SELECT 
  id,
  result_id, 
  ( try1 ).status,
  ( try1 ).no1.scores,
  ( try1 ).no1.duration,
  ( try1 ).no2.scores,
  ( try1 ).no2.duration,
  ( try1 ).no3.scores,
  ( try1 ).no3.duration,
  ( try1 ).no4.scores,
  ( try1 ).no4.duration,
  ( try1 ).no5.scores,
  ( try1 ).no5.duration,
  ( try1 ).no6.scores,
  ( try1 ).no6.duration,
  ( try1 ).checkmarks,
  created_at,
  updated_at
FROM stage1_results
WHERE id = $1;

-- (all role)
-- name: GetStage1try1RelationByResultId :one
SELECT 
  id, 
  result_id
FROM stage1_results
WHERE result_id = $1;

-- name: GetStage1try1Status :one
SELECT 
  ( try1 ).status
FROM stage1_results
WHERE id = $1;

-- (scorer role)
-- name: UpdateStage1try1Checkmarks :one
UPDATE stage1_results
SET 
  try1.checkmarks = $2, 
  updated_at = NOW()
WHERE id = $1
RETURNING ( try1 ).checkmarks;

-- (scorer role)
-- name: UpdateStage1try1NextNo :exec
UPDATE stage1_results
SET try1.status = $2, updated_at = NOW()
WHERE id = $1;

-- (scorer role)
-- name: UpdateStage1FinishSuccess :exec
WITH updated_stage1 AS (
  UPDATE stage1_results
  SET try1.status = '7', shooter_sign = $2, scorer_sign = $3, updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING result_id
)
UPDATE results 
SET stage = '2', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1);

-- (scorer role)
-- name: UpdateStage1FinishFailed :exec
WITH updated_stage1 AS (
  UPDATE stage1_results
  SET try1.status = '7', shooter_sign = $2, scorer_sign = $3, updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING result_id
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1);

-- (scorer role)
-- name: UpdateStage1try1No1 :one
UPDATE stage1_results
SET try1.no1.scores = $2, try1.no1.duration = $3, updated_at = NOW()
WHERE id = $1 
RETURNING ( try1 ).no1.scores, ( try1 ).no1.duration;

-- (scorer role)
-- name: UpdateStage1try1No2 :one
UPDATE stage1_results
SET try1.no2.scores = $2, try1.no2.duration = $3, updated_at = NOW()
WHERE id = $1 
RETURNING ( try1 ).no2.scores, ( try1 ).no2.duration;

-- (scorer role)
-- name: UpdateStage1try1No3 :one
UPDATE stage1_results
SET try1.no3.scores = $2, try1.no3.duration = $3, updated_at = NOW()
WHERE id = $1 
RETURNING ( try1 ).no3.scores, ( try1 ).no3.duration;

-- (scorer role)
-- name: UpdateStage1try1No4 :one
UPDATE stage1_results
SET try1.no4.scores = $2, try1.no4.duration = $3, updated_at = NOW()
WHERE id = $1 
RETURNING ( try1 ).no4.scores, ( try1 ).no4.duration;

-- (scorer role)
-- name: UpdateStage1try1No5 :one
UPDATE stage1_results
SET try1.no5.scores = $2, try1.no5.duration = $3, updated_at = NOW()
WHERE id = $1 
RETURNING ( try1 ).no5.scores, ( try1 ).no5.duration;

-- (scorer role)
-- name: UpdateStage1try1No6 :one
UPDATE stage1_results
SET try1.no6.scores = $2, try1.no6.duration = $3, updated_at = NOW()
WHERE id = $1 
RETURNING ( try1 ).no6.scores, ( try1 ).no6.duration; 

-- (scorer role)
-- name: UpdateStage1try1 :one 
UPDATE stage1_results
SET 
  try1.status = $2,
  try1.no1.scores = $3, try1.no1.duration = $4,
  try1.no2.scores = $5, try1.no2.duration = $6,
  try1.no3.scores = $7, try1.no3.duration = $8,
  try1.no4.scores = $9, try1.no4.duration = $10,
  try1.no5.scores = $11, try1.no5.duration = $12,
  try1.no6.scores = $13, try1.no6.duration = $14,
  try1.checkmarks = $15,
  updated_at = NOW()
WHERE id = $1 
RETURNING 
  ( try1 ).status,
  ( try1 ).no1.scores,
  ( try1 ).no1.duration,
  ( try1 ).no2.scores,
  ( try1 ).no2.duration,
  ( try1 ).no3.scores,
  ( try1 ).no3.duration,
  ( try1 ).no4.scores,
  ( try1 ).no4.duration,
  ( try1 ).no5.scores,
  ( try1 ).no5.duration,
  ( try1 ).no6.scores,
  ( try1 ).no6.duration,
  ( try1 ).checkmarks,
  created_at,
  updated_at;

-- (admin-super role) 
-- name: UpdateStage1Signs :one
UPDATE stage1_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at;

-- (admin-super role) 
-- name: DeleteStage1 :exec
WITH deleted_stage1 AS (
  DELETE FROM stage1_results
  WHERE stage1_results.id = $1
  RETURNING result_id
)
UPDATE results 
SET stage = '1', updated_at = NOW()
WHERE id = (SELECT result_id FROM deleted_stage1);
