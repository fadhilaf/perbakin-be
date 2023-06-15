-- name: CreateStage2 :one
WITH added_stage2_tries AS (
  INSERT INTO stage2_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), added_stage2_results AS (
  INSERT INTO stage2_results (result_id, try1_id)
  SELECT $1, id FROM added_stage2_tries
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
FROM added_stage2_tries, added_stage2_results;

