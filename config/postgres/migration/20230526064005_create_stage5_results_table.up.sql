CREATE TYPE stage5_numbers AS (
  score_a stage123456_scores,
  score_b stage123456_scores,
  score_c stage123456_scores,
  duration stage123456_durations
);

CREATE TYPE stage5_checkmarks AS (
  check1 boolean,
  check2 boolean
);

CREATE TYPE stage5_status AS ENUM ('1', '2', '3');

CREATE TABLE IF NOT EXISTS stage5_tries (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  status stage5_status NOT NULL DEFAULT '1',
  no1 stage5_numbers NOT NULL DEFAULT (ROW(0, 0, 0), ROW(0, 0, 0), ROW(0, 0, 0), ROW(0, 0, 0)),
  no2 stage5_numbers NOT NULL DEFAULT (ROW(0, 0, 0), ROW(0, 0, 0), ROW(0, 0, 0), ROW(0, 0, 0)),
  checkmarks stage5_checkmarks NOT NULL DEFAULT (ROW(false, false))
);

CREATE TABLE IF NOT EXISTS stage5_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  try1_id uuid NOT NULL UNIQUE,
  try2_id uuid DEFAULT NULL UNIQUE,
  is_try2 boolean NOT NULL DEFAULT FALSE,
  shooter_sign varchar(255) DEFAULT NULL,
  scorer_sign varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE,
  CONSTRAINT try1_id 
    FOREIGN KEY (try1_id) 
      REFERENCES stage5_tries (id) ON DELETE CASCADE,
  CONSTRAINT try2_id 
    FOREIGN KEY (try2_id) 
      REFERENCES stage5_tries (id) ON DELETE SET NULL
);
