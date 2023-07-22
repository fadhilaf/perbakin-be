-- name: CreateStage2try2 :one
WITH added_stage2_try2 AS (
  INSERT INTO stage2_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), updated_stage2_results AS (
  UPDATE stage2_results
  SET 
    try2_id = (SELECT id FROM added_stage2_try2),
    is_try2 = true,
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id, is_try2
), updated_stage2_try1 AS (
  UPDATE stage2_tries
  SET status = '4'
  FROM updated_stage2_results
  WHERE id = updated_stage2_results.try1_id
)
SELECT 
  is_try2,
  status,
  no1,
  no2,
  no3,
  checkmarks
FROM added_stage2_try2, updated_stage2_results;

-- (all role) 
-- name: GetStage2try2Status :one
SELECT 
  status
FROM stage2_results
INNER JOIN stage2_tries ON stage2_tries.id = stage2_results.try2_id
WHERE stage2_results.id = $1;

-- (scorer role) 
-- name: UpdateStage2try2Checkmarks :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
SET 
  checkmarks = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try2_id
RETURNING checkmarks;

-- (scorer role) 
-- name: UpdateStage2try2NextNo :exec 
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
  SET status = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try2_id;


-- (scorer role)
-- name: UpdateStage2try2FinishSuccess :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try2_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '4'
  FROM updated_stage2_results
  WHERE id = updated_stage2_results.try2_id
)
UPDATE results 
SET stage = '3', updated_at = NOW()
FROM updated_stage2_results
WHERE id = updated_stage2_results.result_id;

-- (scorer role)
-- name: UpdateStage2try2FinishFailed :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try2_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '4'
  FROM updated_stage2_results
  WHERE id = updated_stage2_results.try2_id
)
UPDATE results 
SET failed = true, updated_at = NOW()
FROM updated_stage2_results
WHERE id = updated_stage2_results.result_id;

-- (scorer role)
-- name: UpdateStage2try2No1 :one 
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
SET no1 = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try2_id
RETURNING no1;

-- (scorer role) 
-- name: UpdateStage2try2No2 :one 
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
SET no2 = $2
FROM updated_stage2_results
WHERE id = updated_stage2_results.try2_id
RETURNING no2;

-- (scorer role) 
-- name: UpdateStage2try2No3 :one 
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
) 
UPDATE stage2_tries 
SET no3 = $2 
FROM updated_stage2_results
WHERE id = updated_stage2_results.try2_id
RETURNING no3;

-- (admin-super role)
-- name: DeleteStage2try2 :exec
WITH deleted_stage2_try2 AS (
  DELETE FROM stage2_tries
  WHERE stage2_tries.id = (SELECT try2_id FROM stage2_results WHERE stage2_results.id = $1)
), updated_stage2_results AS (
  UPDATE stage2_results
  SET
    try2_id = NULL,
    is_try2 = false,
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET status = '3'
FROM updated_stage2_results
WHERE stage2_tries.id = updated_stage2_results.try1_id;

-- (admin-super role)
-- name: UpdateStage2try2 :one 
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
), updated_stage2_try2 AS (
  UPDATE stage2_tries
  SET 
    status = sqlc.arg(try2_status),
    no1 = sqlc.arg(try2_no1),
    no2 = sqlc.arg(try2_no2),
    no3 = sqlc.arg(try2_no3),
    checkmarks = sqlc.arg(try2_checkmarks)
  WHERE id IN (SELECT try2_id FROM updated_stage2_results WHERE try2_id IS NOT NULL)
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
  updated_stage2_try2.status AS try2_status,
  updated_stage2_try2.no1 AS try2_no1,
  updated_stage2_try2.no2 AS try2_no2,
  updated_stage2_try2.no3 AS try2_no3,
  updated_stage2_try2.checkmarks AS try2_checkmarks,
  updated_at
FROM updated_stage2_results, updated_stage2_try1, updated_stage2_try2;
