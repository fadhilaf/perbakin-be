CREATE TYPE stage246_checkmarks AS (
  check1 boolean,
  check2 boolean,
  check3 boolean
);

CREATE TYPE stage246_status AS ENUM ('1', '2', '3', '4');

CREATE TABLE IF NOT EXISTS stage2_tries (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  status stage246_status NOT NULL DEFAULT '1',
  no1 stage123_numbers NOT NULL DEFAULT (ROW(0, 0, 0), ROW(0, 0, 0)),
  no2 stage123_numbers NOT NULL DEFAULT (ROW(0, 0, 0), ROW(0, 0, 0)),
  no3 stage123_numbers NOT NULL DEFAULT (ROW(0, 0, 0), ROW(0, 0, 0)),
  checkmarks stage246_checkmarks NOT NULL DEFAULT (ROW(false, false, false))
);

CREATE TABLE IF NOT EXISTS stage2_results (
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
      REFERENCES stage2_tries (id) ON DELETE CASCADE,
  CONSTRAINT try2_id 
    FOREIGN KEY (try2_id) 
      REFERENCES stage2_tries (id) ON DELETE SET NULL
);
