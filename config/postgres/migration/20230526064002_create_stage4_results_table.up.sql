CREATE TYPE stage46_score_types AS (
  a stage123456_scores,
  b stage123456_scores
);

CREATE TYPE stage46_numbers AS (
  score_types stage46_score_types,
  duration stage123456_durations
);

CREATE TYPE stage46_tries AS (
  status stage246_status,
  no1 stage46_numbers,
  no2 stage46_numbers,
  no3 stage46_numbers,
  checkmarks stage246_checkmarks
);

CREATE TABLE IF NOT EXISTS stage4_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  try1 stage46_tries NOT NULL DEFAULT ROW( 
    '1',
    ROW(
      ROW(
        ROW(0,0,0),
        ROW(0,0,0)
      ),
      ROW(0,0,0)
    ),
    ROW(
      ROW(
        ROW(0,0,0),
        ROW(0,0,0)
      ),
      ROW(0,0,0)
    ),
    ROW(
      ROW(
        ROW(0,0,0),
        ROW(0,0,0)
      ),
      ROW(0,0,0)
    ),
    ROW(false,false,false)
  ),
  try2 stage46_tries DEFAULT NULL,  
  is_try2 boolean NOT NULL DEFAULT FALSE,
  shooter_sign varchar(255) DEFAULT NULL,
  scorer_sign varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE
);
