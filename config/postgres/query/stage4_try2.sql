-- name: CreateStage4try2 :one
WITH added_stage4_try2 AS (
  INSERT INTO stage46_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), updated_stage4_results AS (
  UPDATE stage4_results
  SET 
    try2_id = (SELECT id FROM added_stage4_try2),
    is_try2 = true,
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id, is_try2
), updated_stage4_try1 AS (
  UPDATE stage46_tries
  SET status = '4'
  WHERE id = (SELECT try1_id FROM updated_stage4_results)
)
SELECT 
  is_try2,
  status,
  no1,
  no2,
  no3,
  checkmarks
FROM added_stage4_try2, updated_stage4_results;

-- (all role) 
-- name: GetStage4try2Status :one
SELECT 
  status
FROM stage4_results
INNER JOIN stage46_tries ON stage46_tries.id = stage4_results.try2_id
WHERE stage4_results.id = $1;

-- (scorer role) 
-- name: UpdateStage4try2Checkmarks :one
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try2_id
)
UPDATE stage46_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try2_id FROM stage4_results)
RETURNING checkmarks;

-- (scorer role) 
-- name: UpdateStage4try2NextNo :exec 
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try2_id
)
UPDATE stage46_tries
  SET status = $2
WHERE id = (SELECT try2_id FROM stage4_results);


-- (scorer role)
-- name: UpdateStage4try2FinishSuccess :exec
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage4_results.id = $1
  RETURNING result_id, try2_id
), updated_stage4_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  WHERE id = (SELECT try2_id FROM updated_stage4_results)
)
UPDATE results 
SET stage = '5', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage4_results);

-- (scorer role)
-- name: UpdateStage4try2FinishFailed :exec
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage4_results.id = $1
  RETURNING result_id, try2_id
), updated_stage4_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  WHERE id = (SELECT try2_id FROM updated_stage4_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage4_results);

-- (scorer role)
-- name: UpdateStage4try2No1 :one 
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try2_id
)
UPDATE stage46_tries
SET no1 = $2
WHERE id = (SELECT try2_id FROM stage4_results)
RETURNING no1;

-- (scorer role) 
-- name: UpdateStage4try2No2 :one 
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try2_id
)
UPDATE stage46_tries
SET no2 = $2
WHERE id = (SELECT try2_id FROM stage4_results)
RETURNING no2;

-- (scorer role) 
-- name: UpdateStage4try2No3 :one 
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try2_id
) 
UPDATE stage46_tries 
SET no3 = $2 
WHERE id = (SELECT try2_id FROM stage4_results)
RETURNING no3;

-- (admin-super role)
-- name: DeleteStage4try2 :exec
WITH deleted_stage4_try2 AS (
  DELETE FROM stage46_tries
  WHERE stage46_tries.id = (SELECT try2_id FROM stage4_results WHERE stage4_results.id = $1)
), updated_stage4_results AS (
  UPDATE stage4_results
  SET
    try2_id = NULL,
    is_try2 = false,
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET status = '6'
WHERE stage46_tries.id = (SELECT try1_id FROM updated_stage4_results);

-- (admin-super role)
-- name: UpdateStage4try2 :one 
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage4_try1 AS (
  UPDATE stage46_tries
  SET 
    status = sqlc.arg(try1_status),
    no1 = sqlc.arg(try1_no1),
    no2 = sqlc.arg(try1_no2),
    no3 = sqlc.arg(try1_no3),
    checkmarks = sqlc.arg(try1_checkmarks)
  WHERE id = (SELECT try1_id FROM stage4_results)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
), updated_stage4_try2 AS (
  UPDATE stage46_tries
  SET 
    status = sqlc.arg(try2_status),
    no1 = sqlc.arg(try2_no1),
    no2 = sqlc.arg(try2_no2),
    no3 = sqlc.arg(try2_no3),
    checkmarks = sqlc.arg(try2_checkmarks)
  WHERE id = (SELECT try2_id FROM stage4_results WHERE try2_id IS NOT NULL)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
)
SELECT 
  updated_stage4_try1.status AS try1_status,
  updated_stage4_try1.no1 AS try1_no1,
  updated_stage4_try1.no2 AS try1_no2,
  updated_stage4_try1.no3 AS try1_no3,
  updated_stage4_try1.checkmarks AS try1_checkmarks,
  updated_stage4_try2.status AS try2_status,
  updated_stage4_try2.no1 AS try2_no1,
  updated_stage4_try2.no2 AS try2_no2,
  updated_stage4_try2.no3 AS try2_no3,
  updated_stage4_try2.checkmarks AS try2_checkmarks,
  updated_at
FROM updated_stage4_results, updated_stage4_try1, updated_stage4_try2;
