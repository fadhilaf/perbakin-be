-- name: CreateStage6 :one
WITH added_stage6_try1 AS (
  INSERT INTO stage46_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), added_stage6_results AS (
  INSERT INTO stage6_results (result_id, try1_id)
  SELECT $1, id FROM added_stage6_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage6_results.id, 
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
FROM added_stage6_try1, added_stage6_results;

-- (all role)
-- name: GetStage6ById :one
SELECT 
  stage6_results.id,
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
FROM stage6_results
INNER JOIN stage46_tries AS try1 ON try1.id = stage6_results.try1_id
LEFT JOIN stage46_tries AS try2 ON try2.id = stage6_results.try2_id
WHERE stage6_results.id = $1;

-- (all role)
-- name: GetStage6RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage6_results
WHERE result_id = $1;

-- (all role)
-- name: GetStage6try2ExistById :one
SELECT 
  is_try2
FROM stage6_results
WHERE id = $1;

-- name: GetStage6try1Status :one
SELECT 
  status
FROM stage6_results
INNER JOIN stage46_tries ON stage46_tries.id = stage6_results.try1_id
WHERE stage6_results.id = $1;

-- (scorer role)
-- name: UpdateStage6try1Checkmarks :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try1_id FROM stage6_results)
RETURNING checkmarks;

-- (scorer role)
-- name: UpdateStage6try1NextNo :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
  SET status = $2
WHERE id = (SELECT try1_id FROM stage6_results);

-- (scorer role)
-- name: UpdateStage6try1FinishSuccess :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage6_results.id = $1
  RETURNING result_id, try1_id
), updated_stage6_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  WHERE id = (SELECT try1_id FROM updated_stage6_results)
)
UPDATE results 
SET stage = '7', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage6_results);

-- (scorer role)
-- name: UpdateStage6try1FinishFailed :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage6_results.id = $1
  RETURNING result_id, try1_id
), updated_stage6_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  WHERE id = (SELECT try1_id FROM updated_stage6_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage6_results);

-- (scorer role)
-- name: UpdateStage6NextTry :exec 
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET status = '4'
WHERE id = (SELECT try1_id FROM stage6_results);

-- (scorer role)
-- name: UpdateStage6try1No1 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET no1 = $2
WHERE id = (SELECT try1_id FROM stage6_results)
RETURNING no1;

-- (scorer role)
-- name: UpdateStage6try1No2 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries 
SET no2 = $2
WHERE id = (SELECT try1_id FROM stage6_results)
RETURNING no2;

-- (scorer role)
-- name: UpdateStage6try1No3 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET no3 = $2
WHERE id = (SELECT try1_id FROM stage6_results) 
RETURNING no3; 

-- (admin-super role) 
-- name: UpdateStage6Signs :one
UPDATE stage6_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at;

-- (admin-super role)
-- name: UpdateStage6try1 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage6_try1 AS (
  UPDATE stage46_tries
  SET 
    status = sqlc.arg(try1_status),
    no1 = sqlc.arg(try1_no1),
    no2 = sqlc.arg(try1_no2),
    no3 = sqlc.arg(try1_no3),
    checkmarks = sqlc.arg(try1_checkmarks)
  WHERE id = (SELECT try1_id FROM stage6_results)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
)
SELECT 
  updated_stage6_try1.status AS try1_status,
  updated_stage6_try1.no1 AS try1_no1,
  updated_stage6_try1.no2 AS try1_no2,
  updated_stage6_try1.no3 AS try1_no3,
  updated_stage6_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage6_results, updated_stage6_try1;

-- (admin-super role)
-- name: FinishStage6 :exec 
WITH get_stage6 AS (
  SELECT 
    result_id, try1_id, try2_id
  FROM stage6_results
  WHERE stage6_results.id = $1
), updated_stage6try1 AS (
  UPDATE stage46_tries
  SET status = '4'
  WHERE id = (SELECT try1_id FROM get_stage6)
), updated_stage6try2 AS (
  UPDATE stage46_tries
  SET status = '4'
  WHERE id = (SELECT try2_id FROM get_stage6 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '7', updated_at = NOW()
WHERE id = (SELECT result_id FROM get_stage6);

-- (admin-super role) 
-- name: DeleteStage6 :exec
WITH deleted_stage6 AS (
  DELETE FROM stage6_results
  WHERE stage6_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage6try1 AS (
  DELETE FROM stage46_tries
  WHERE stage46_tries.id = (SELECT try1_id FROM deleted_stage6)
)
DELETE FROM stage46_tries
WHERE stage46_tries.id = (SELECT try2_id FROM deleted_stage6 WHERE try2_id IS NOT NULL);
