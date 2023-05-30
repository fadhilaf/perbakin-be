-- name: CreateStage0try1 :one
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
  ( try1 ).checkmarks;
