-- name: CreateStage0try2 :one
UPDATE stage1_results
SET 
  is_try2 = true,
  try2 = ROW( 
    '1',
    ROW(
      ROW(0,0,0),
      '00:00:00'
    ),
    ROW(
      ROW(0,0,0),
      '00:00:00'
    ),
    ROW(
      ROW(0,0,0),
      '00:00:00'
    ),
    ROW(
      ROW(0,0,0),
      '00:00:00'
    ),
    ROW(
      ROW(0,0,0),
      '00:00:00'
    ),
    ROW(
      ROW(0,0,0),
      '00:00:00'
    ),
    ROW(false,false,false,false,false,false)
  )
WHERE result_id = $1
RETURNING 
  id, 
  result_id, 
  ( try2 ).status,
  ( try2 ).no1.scores,
  ( try2 ).no1.duration,
  ( try2 ).no2.scores,
  ( try2 ).no2.duration,
  ( try2 ).no3.scores,
  ( try2 ).no3.duration,
  ( try2 ).no4.scores,
  ( try2 ).no4.duration,
  ( try2 ).no5.scores,
  ( try2 ).no5.duration,
  ( try2 ).no6.scores,
  ( try2 ).no6.duration,
  ( try2 ).checkmarks;
