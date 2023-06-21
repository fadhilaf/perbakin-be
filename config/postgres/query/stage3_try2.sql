-- name: CreateStage3try2 :one
WITH added_stage3_try2 AS (
  INSERT INTO stage13_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, no4, no5, no6, checkmarks
), updated_stage3_results AS (
  UPDATE stage3_results
  SET 
    try2_id = (SELECT id FROM added_stage3_try2),
    is_try2 = true,
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try1_id, is_try2
), updated_stage3_try1 AS (
  UPDATE stage13_tries
  SET status = '7'
  WHERE id = (SELECT try1_id FROM updated_stage3_results)
)
SELECT 
  is_try2,
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks
FROM added_stage3_try2, updated_stage3_results;

-- (all role) 
-- name: GetStage3try2Status :one
SELECT 
  status
FROM stage3_results
INNER JOIN stage13_tries ON stage13_tries.id = stage3_results.try2_id
WHERE stage3_results.id = $1;

-- (scorer role) 
-- name: UpdateStage3try2Checkmarks :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING checkmarks;

-- (scorer role) 
-- name: UpdateStage3try2NextNo :exec 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
  SET status = $2
WHERE id = (SELECT try2_id FROM stage3_results);


-- (scorer role)
-- name: UpdateStage3try2FinishSuccess :exec
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage3_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage3_results)
)
UPDATE results 
SET stage = '2', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage3_results);

-- (scorer role)
-- name: UpdateStage3try2FinishFailed :exec
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage3_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage3_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage3_results);

-- (scorer role)
-- name: UpdateStage3try2No1 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no1 = $2
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING no1;

-- (scorer role) 
-- name: UpdateStage3try2No2 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no2 = $2
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING no2;

-- (scorer role) 
-- name: UpdateStage3try2No3 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
) 
UPDATE stage13_tries 
SET no3 = $2 
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING no3;

-- (scorer role) 
-- name: UpdateStage3try2No4 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no4 = $2
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING no4;

-- (scorer role) 
-- name: UpdateStage3try2No5 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no5 = $2
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING no5;

-- (scorer role) 
-- name: UpdateStage3try2No6 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no6 = $2
WHERE id = (SELECT try2_id FROM stage3_results)
RETURNING no6;

-- (admin-super role)
-- name: DeleteStage3try2 :exec
WITH deleted_stage3_try2 AS (
  DELETE FROM stage13_tries
  WHERE stage13_tries.id = (SELECT try2_id FROM stage3_results WHERE stage3_results.id = $1)
), updated_stage3_results AS (
  UPDATE stage3_results
  SET
    try2_id = NULL,
    is_try2 = false,
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET status = '6'
WHERE stage13_tries.id = (SELECT try1_id FROM updated_stage3_results);

-- (admin-super role)
-- name: UpdateStage3try2 :one 
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage3_try1 AS (
  UPDATE stage13_tries
  SET 
    status = sqlc.arg(try1_status),
    no1 = sqlc.arg(try1_no1),
    no2 = sqlc.arg(try1_no2),
    no3 = sqlc.arg(try1_no3),
    no4 = sqlc.arg(try1_no4),
    no5 = sqlc.arg(try1_no5),
    no6 = sqlc.arg(try1_no6),
    checkmarks = sqlc.arg(try1_checkmarks)
  WHERE id = (SELECT try1_id FROM stage3_results)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    no4,
    no5,
    no6,
    checkmarks
), updated_stage3_try2 AS (
  UPDATE stage13_tries
  SET 
    status = sqlc.arg(try2_status),
    no1 = sqlc.arg(try2_no1),
    no2 = sqlc.arg(try2_no2),
    no3 = sqlc.arg(try2_no3),
    no4 = sqlc.arg(try2_no4),
    no5 = sqlc.arg(try2_no5),
    no6 = sqlc.arg(try2_no6),
    checkmarks = sqlc.arg(try2_checkmarks)
  WHERE id = (SELECT try2_id FROM stage3_results WHERE try2_id IS NOT NULL)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    no4,
    no5,
    no6,
    checkmarks
)
SELECT 
  updated_stage3_try1.status AS try1_status,
  updated_stage3_try1.no1 AS try1_no1,
  updated_stage3_try1.no2 AS try1_no2,
  updated_stage3_try1.no3 AS try1_no3,
  updated_stage3_try1.no4 AS try1_no4,
  updated_stage3_try1.no5 AS try1_no5,
  updated_stage3_try1.no6 AS try1_no6,
  updated_stage3_try1.checkmarks AS try1_checkmarks,
  updated_stage3_try2.status AS try2_status,
  updated_stage3_try2.no1 AS try2_no1,
  updated_stage3_try2.no2 AS try2_no2,
  updated_stage3_try2.no3 AS try2_no3,
  updated_stage3_try2.no4 AS try2_no4,
  updated_stage3_try2.no5 AS try2_no5,
  updated_stage3_try2.no6 AS try2_no6,
  updated_stage3_try2.checkmarks AS try2_checkmarks,
  updated_at
FROM updated_stage3_results, updated_stage3_try1, updated_stage3_try2;
