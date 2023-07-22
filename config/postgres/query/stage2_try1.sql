-- name: CreateStage2 :one
WITH added_stage2_try1 AS (
  INSERT INTO stage2_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), added_stage2_results AS (
  INSERT INTO stage2_results (result_id, try1_id)
  SELECT $1, id FROM added_stage2_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage2_results.id, 
  result_id, 
  status,
  no1,
  no2,
  no3,
  checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM added_stage2_try1, added_stage2_results;

-- (all role)
-- name: GetStage2ById :one
SELECT 
  stage2_results.id,
  result_id, 
  try1.status AS try1_status,
  try1.no1 AS try1_no1,
  try1.no2 AS try1_no2,
  try1.no3 AS try1_no3,
  try1.checkmarks AS try1_checkmarks,
  try2.status AS try2_status,
  try2.no1 AS try2_no1,
  try2.no2 AS try2_no2,
  try2.no3 AS try2_no3,
  try2.checkmarks AS try2_checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM stage2_results
INNER JOIN stage2_tries AS try1 ON try1.id = stage2_results.try1_id
LEFT JOIN stage2_tries AS try2 ON try2.id = stage2_results.try2_id
WHERE stage2_results.id = $1;

-- (all role)
-- name: GetStage2RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage2_results
WHERE result_id = $1;

-- (all role)
-- name: GetStage2try2ExistById :one
SELECT 
  is_try2
FROM stage2_results
WHERE id = $1;

-- name: GetStage2try1Status :one
SELECT 
  status
FROM stage2_results
INNER JOIN stage2_tries ON stage2_tries.id = stage2_results.try1_id
WHERE stage2_results.id = $1;

-- (scorer role)
-- name: UpdateStage2try1Checkmarks :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET 
  checkmarks = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try1_id
RETURNING checkmarks;

-- (scorer role)
-- name: UpdateStage2try1NextNo :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
  SET status = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try1_id;

-- (scorer role)
-- name: UpdateStage2try1FinishSuccess :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try1_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '4'
  FROM updated_stage2_results
  WHERE id = updated_stage2_results.try1_id
)
UPDATE results 
SET stage = '3', updated_at = NOW()
FROM updated_stage2_results
WHERE id = updated_stage2_results.result_id;

-- (scorer role)
-- name: UpdateStage2try1FinishFailed :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try1_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '3'
  FROM updated_stage2_results
  WHERE id = updated_stage2_results.try1_id
)
UPDATE results 
SET failed = true, updated_at = NOW()
FROM updated_stage2_results
WHERE id = updated_stage2_results.result_id;

-- (scorer role)
-- name: UpdateStage2NextTry :exec 
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET status = '4'
FROM updated_stage2_results
WHERE id = updated_stage2_results.try1_id;

-- (scorer role)
-- name: UpdateStage2try1No1 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET no1 = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try1_id
RETURNING no1;

-- (scorer role)
-- name: UpdateStage2try1No2 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries 
SET no2 = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try1_id
RETURNING no2;

-- (scorer role)
-- name: UpdateStage2try1No3 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET no3 = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try1_id 
RETURNING no3; 

-- (admin-super role) 
-- name: UpdateStage2Signs :one
UPDATE stage2_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at;

-- (admin-super role)
-- name: UpdateStage2try1 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage2_try1 AS (
  UPDATE stage2_tries
  SET 
    status = sqlc.arg(try1_status),
    no1 = sqlc.arg(try1_no1),
    no2 = sqlc.arg(try1_no2),
    no3 = sqlc.arg(try1_no3),
    checkmarks = sqlc.arg(try1_checkmarks)
  FROM updated_stage2_results
  WHERE id = updated_stage2_results.try1_id 
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
)
SELECT 
  updated_stage2_try1.status AS try1_status,
  updated_stage2_try1.no1 AS try1_no1,
  updated_stage2_try1.no2 AS try1_no2,
  updated_stage2_try1.no3 AS try1_no3,
  updated_stage2_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage2_results, updated_stage2_try1;

-- (admin-super role)
-- name: FinishStage2 :exec 
WITH get_stage2 AS (
  SELECT 
    result_id, try1_id, try2_id
  FROM stage2_results
  WHERE stage2_results.id = $1
), updated_stage2try1 AS (
  UPDATE stage2_tries
  SET status = '4'
  WHERE id = (SELECT try1_id FROM get_stage2)
), updated_stage2try2 AS (
  UPDATE stage2_tries
  SET status = '4'
  WHERE id = (SELECT try2_id FROM get_stage2 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '3', updated_at = NOW()
WHERE id = (SELECT result_id FROM get_stage2);

-- (admin-super role) 
-- name: DeleteStage2 :exec
WITH deleted_stage2 AS (
  DELETE FROM stage2_results
  WHERE stage2_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage2try1 AS (
  DELETE FROM stage2_tries
  WHERE stage2_tries.id IN (SELECT try1_id FROM deleted_stage2)
)
DELETE FROM stage2_tries
WHERE stage2_tries.id = (SELECT try2_id FROM deleted_stage2 WHERE try2_id IS NOT NULL);
