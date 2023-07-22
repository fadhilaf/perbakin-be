-- name: CreateStage5 :one
WITH added_stage5_try1 AS (
  INSERT INTO stage5_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, checkmarks
), added_stage5_results AS (
  INSERT INTO stage5_results (result_id, try1_id)
  SELECT $1, id FROM added_stage5_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage5_results.id, 
  result_id, 
  status,
  no1,
  no2,
  checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM added_stage5_try1, added_stage5_results;

-- (all role)
-- name: GetStage5ById :one
SELECT 
  stage5_results.id,
  result_id, 
  try1.status AS try1_status,
  try1.no1 AS try1_no1,
  try1.no2 AS try1_no2,
  try1.checkmarks AS try1_checkmarks,
  try2.status AS try2_status,
  try2.no1 AS try2_no1,
  try2.no2 AS try2_no2,
  try2.checkmarks AS try2_checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM stage5_results
INNER JOIN stage5_tries AS try1 ON try1.id = stage5_results.try1_id
LEFT JOIN stage5_tries AS try2 ON try2.id = stage5_results.try2_id
WHERE stage5_results.id = $1;

-- (all role)
-- name: GetStage5RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage5_results
WHERE result_id = $1;

-- (all role)
-- name: GetStage5try2ExistById :one
SELECT 
  is_try2
FROM stage5_results
WHERE id = $1;

-- name: GetStage5try1Status :one
SELECT 
  status
FROM stage5_results
INNER JOIN stage5_tries ON stage5_tries.id = stage5_results.try1_id
WHERE stage5_results.id = $1;

-- (scorer role)
-- name: UpdateStage5try1Checkmarks :one
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW()
  WHERE stage5_results.id = $1
  RETURNING try1_id
)
UPDATE stage5_tries
SET 
  checkmarks = $2
FROM updated_stage5_results
WHERE id = updated_stage5_results.try1_id
RETURNING checkmarks;

-- (scorer role)
-- name: UpdateStage5try1NextNo :exec
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW()
  WHERE stage5_results.id = $1
  RETURNING try1_id
)
UPDATE stage5_tries
  SET status = $2
FROM updated_stage5_results
WHERE id = updated_stage5_results.try1_id;

-- (scorer role)
-- name: UpdateStage5try1FinishSuccess :exec
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage5_results.id = $1
  RETURNING result_id, try1_id
), updated_stage5_tries AS (
  UPDATE stage5_tries
    SET status = '3'
  FROM updated_stage5_results
  WHERE id = updated_stage5_results.try1_id
)
UPDATE results 
SET stage = '6', updated_at = NOW()
FROM updated_stage5_results
WHERE id = updated_stage5_results.result_id;

-- (scorer role)
-- name: UpdateStage5try1FinishFailed :exec
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage5_results.id = $1
  RETURNING result_id, try1_id
), updated_stage5_tries AS (
  UPDATE stage5_tries
    SET status = '3'
  FROM updated_stage5_results
  WHERE id = updated_stage5_results.try1_id
)
UPDATE results 
SET failed = true, updated_at = NOW()
FROM updated_stage5_results
WHERE id = updated_stage5_results.result_id;

-- (scorer role)
-- name: UpdateStage5NextTry :exec 
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage5_results.id = $1
  RETURNING try1_id
)
UPDATE stage5_tries
SET status = '3'
FROM updated_stage5_results
WHERE id = updated_stage5_results.try1_id;

-- (scorer role)
-- name: UpdateStage5try1No1 :one
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW()
  WHERE stage5_results.id = $1
  RETURNING try1_id
)
UPDATE stage5_tries
SET no1 = $2
FROM updated_stage5_results
WHERE id = updated_stage5_results.try1_id
RETURNING no1;

-- (scorer role)
-- name: UpdateStage5try1No2 :one
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW()
  WHERE stage5_results.id = $1
  RETURNING try1_id
)
UPDATE stage5_tries 
SET no2 = $2
FROM updated_stage5_results
WHERE id = updated_stage5_results.try1_id
RETURNING no2;

-- (admin-super role) 
-- name: UpdateStage5Signs :one
UPDATE stage5_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at;

-- (admin-super role)
-- name: UpdateStage5try1 :one
WITH updated_stage5_results AS (
  UPDATE stage5_results
  SET
    updated_at = NOW()
  WHERE stage5_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage5_try1 AS (
  UPDATE stage5_tries
  SET 
    status = sqlc.arg(try1_status),
    no1 = sqlc.arg(try1_no1),
    no2 = sqlc.arg(try1_no2),
    checkmarks = sqlc.arg(try1_checkmarks)
  FROM updated_stage5_results
  WHERE id = updated_stage5_results.try1_id
  RETURNING 
    status,
    no1,
    no2,
    checkmarks
)
SELECT 
  updated_stage5_try1.status AS try1_status,
  updated_stage5_try1.no1 AS try1_no1,
  updated_stage5_try1.no2 AS try1_no2,
  updated_stage5_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage5_results, updated_stage5_try1;

-- (admin-super role)
-- name: FinishStage5 :exec 
WITH get_stage5 AS (
  SELECT 
    result_id, try1_id, try2_id
  FROM stage5_results
  WHERE stage5_results.id = $1
), updated_stage5try1 AS (
  UPDATE stage5_tries
  SET status = '3'
  FROM get_stage5
  WHERE id = get_stage5.try1_id
), updated_stage5try2 AS (
  UPDATE stage5_tries
  SET status = '3'
  WHERE id IN (SELECT try2_id FROM get_stage5 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '6', updated_at = NOW()
FROM get_stage5
WHERE id = get_stage5.result_id;

-- (admin-super role) 
-- name: DeleteStage5 :exec
WITH deleted_stage5 AS (
  DELETE FROM stage5_results
  WHERE stage5_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage5try1 AS (
  DELETE FROM stage5_tries
  WHERE stage5_tries.id IN (SELECT try1_id FROM deleted_stage5)
)
DELETE FROM stage5_tries
WHERE stage5_tries.id = (SELECT try2_id FROM deleted_stage5 WHERE try2_id IS NOT NULL);
