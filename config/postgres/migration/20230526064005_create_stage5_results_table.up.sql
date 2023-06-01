CREATE TYPE stage5_score_types AS (
  a stage123456_scores,
  b stage123456_scores,
  c stage123456_scores
);

CREATE TYPE stage5_numbers AS (
  score_types stage5_score_types,
  duration stage123456_durations
);

CREATE TYPE stage5_checkmarks AS (
  check1 boolean,
  check2 boolean
);

CREATE TYPE stage5_status AS ENUM ('1', '2', '3');

CREATE TYPE stage5_tries AS (
  status stage5_status,
  no1 stage5_numbers,
  no2 stage5_numbers,
  checkmarks stage5_checkmarks
);

CREATE TABLE IF NOT EXISTS stage5_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  try1 stage5_tries NOT NULL DEFAULT ROW( 
    '1',
    ROW(
      ROW(
        ROW(0,0,0),
        ROW(0,0,0),
        ROW(0,0,0)
      ),
      ROW(0,0,0)
    ),
    ROW(
      ROW(
        ROW(0,0,0),
        ROW(0,0,0),
        ROW(0,0,0)
      ),
      ROW(0,0,0)
    ),
    ROW(false,false)
  ),
  try2 stage5_tries DEFAULT NULL,
  is_try2 boolean NOT NULL DEFAULT FALSE,
  shooter_sign varchar(255) DEFAULT NULL,
  scorer_sign varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE
);
