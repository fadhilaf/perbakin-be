-- name: CreateStage1 :one
WITH added_stage13_tries AS (
  INSERT INTO stage13_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, no4, no5, no6, checkmarks
), added_stage1_results AS (
  INSERT INTO stage1_results (result_id, try1_id)
  SELECT $1, id FROM added_stage13_tries
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage1_results.id, 
  result_id, 
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM added_stage13_tries, added_stage1_results;

-- (all role)
-- name: GetStage1ById :one
SELECT 
  stage1_results.id,
  result_id, 
  try1.status AS try1_status,
  try1.no1 AS try1_no1,
  try1.no2 AS try1_no2,
  try1.no3 AS try1_no3,
  try1.no4 AS try1_no4,
  try1.no5 AS try1_no5,
  try1.no6 AS try1_no6,
  try1.checkmarks AS try1_checkmarks,
  try2.status AS try2_status,
  try2.no1 AS try2_no1,
  try2.no2 AS try2_no2,
  try2.no3 AS try2_no3,
  try2.no4 AS try2_no4,
  try2.no5 AS try2_no5,
  try2.no6 AS try2_no6,
  try2.checkmarks AS try2_checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM stage1_results
INNER JOIN stage13_tries AS try1 ON try1.id = stage1_results.try1_id
LEFT JOIN stage13_tries AS try2 ON try2.id = stage1_results.try2_id
WHERE stage1_results.id = $1;

-- (all role)
-- name: GetStage1RelationByResultId :one
SELECT 
  id, 
  result_id
FROM stage1_results
WHERE result_id = $1;

-- name: GetStage1try1Status :one
SELECT 
  status
FROM stage1_results
INNER JOIN stage13_tries ON stage13_tries.id = stage1_results.try1_id
WHERE stage1_results.id = $1;

-- (scorer role)
-- name: UpdateStage1try1Checkmarks :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING checkmarks;

-- (scorer role)
-- name: UpdateStage1try1NextNo :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
  SET status = $2
WHERE id = (SELECT try1_id FROM stage1_results);

-- (scorer role)
-- name: UpdateStage1try1FinishSuccess :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try1_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try1_id FROM updated_stage1_results)
)
UPDATE results 
SET stage = '2', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results);

-- (scorer role)
-- name: UpdateStage1try1FinishFailed :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try1_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try1_id FROM updated_stage1_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results);

-- (scorer role)
-- name: UpdateStage1NextTry :exec 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET status = '7'
WHERE id = (SELECT try1_id FROM stage1_results);

-- (scorer role)
-- name: UpdateStage1try1No1 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET no1 = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING no1;

-- (scorer role)
-- name: UpdateStage1try1No2 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries 
SET no2 = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING no2;

-- (scorer role)
-- name: UpdateStage1try1No3 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET no3 = $2
WHERE id = (SELECT try1_id FROM stage1_results) 
RETURNING no3; 

-- (scorer role)
-- name: UpdateStage1try1No4 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET no4 = $2
WHERE id = (SELECT try1_id FROM stage1_results) 
RETURNING no4;

-- (scorer role)
-- name: UpdateStage1try1No5 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
) 
UPDATE stage13_tries 
SET no5 = $2
WHERE id = (SELECT try1_id FROM stage1_results) 
RETURNING no5; 

-- (scorer role)
-- name: UpdateStage1try1No6 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1 
  RETURNING try1_id
) 
UPDATE stage13_tries
SET no6 = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING no6;

-- (admin-super role)
-- name: UpdateStage1try1 :one 
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1 
  RETURNING try1_id, updated_at
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
  WHERE id = (SELECT try1_id FROM stage1_results)
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
  RETURNING result_id, try1_id, try2_id
), deleted_stage1try1 AS (
  DELETE FROM stage13_tries
  WHERE stage13_tries.id = (SELECT try1_id FROM deleted_stage1)
), deleted_stage1try2 AS (
  DELETE FROM stage13_tries
  WHERE stage13_tries.id = (SELECT try2_id FROM deleted_stage1 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '1', updated_at = NOW()
WHERE id = (SELECT result_id FROM deleted_stage1);
