-- name: CreateStage1try2 :one
WITH added_stage13_tries AS (
  INSERT INTO stage13_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, no4, no5, no6, checkmarks
), updated_stage1_results AS (
  UPDATE stage1_results
  SET 
    try2_id = (SELECT id FROM added_stage13_tries),
    updated_at = NOW()
  WHERE stage1_results.id = $1
)
SELECT 
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks
FROM added_stage13_tries;

-- (all role) 
-- name: GetStage1try2Status :one
SELECT 
  status
FROM stage1_results
INNER JOIN stage13_tries ON stage13_tries.id = stage1_results.try2_id
WHERE stage1_results.id = $1;

-- (scorer role) 
-- name: UpdateStage1try2Checkmarks :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING checkmarks;

-- (scorer role) 
-- name: UpdateStage1try2NextNo :exec 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
  SET status = $2
WHERE id = (SELECT try2_id FROM stage1_results);


-- (scorer role)
-- name: UpdateStage1try2FinishSuccess :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage1_results)
)
UPDATE results 
SET stage = '2', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results);

-- (scorer role)
-- name: UpdateStage1try2FinishFailed :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage1_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results);

-- (scorer role)
-- name: UpdateStage1try2No1 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no1 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no1;

-- (scorer role) 
-- name: UpdateStage1try2No2 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no2 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no2;

-- (scorer role) 
-- name: UpdateStage1try2No3 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
) 
UPDATE stage13_tries 
SET no3 = $2 
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no3;

-- (scorer role) 
-- name: UpdateStage1try2No4 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no4 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no4;

-- (scorer role) 
-- name: UpdateStage1try2No5 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no5 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no5;

-- (scorer role) 
-- name: UpdateStage1try2No6 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no6 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no6;

-- (admin-super role)
-- name: UpdateStage1try2 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1 
  RETURNING try2_id, updated_at
), updated_stage13_tries AS (
  UPDATE stage13_tries
  SET 
    status = $2, 
    no1 = $3, 
    no2 = $4, 
    no3 = $5, 
    no4 = $6,
    no5 = $7, 
    no6 = $8, 
    checkmarks = $9
  WHERE id = (SELECT try2_id FROM stage1_results)
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
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks,
  updated_at
FROM updated_stage1_results, updated_stage13_tries;
