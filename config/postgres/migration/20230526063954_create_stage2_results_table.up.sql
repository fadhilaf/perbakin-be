CREATE TYPE stage246_checkmarks AS (
  check1 boolean,
  check2 boolean,
  check3 boolean
);

CREATE TYPE stage246_status AS ENUM ('1', '2', '3', '4');

CREATE TYPE stage2_tries AS (
  status stage246_status,
  no1 stage123_numbers,
  no2 stage123_numbers,
  no3 stage123_numbers,
  checkmarks stage246_checkmarks
);

CREATE TABLE IF NOT EXISTS stage2_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  try1 stage2_tries NOT NULL DEFAULT ROW( 
    '1',
    ROW(
      ROW(0,0,0),
      ROW(0,0,0)
    ),
    ROW(
      ROW(0,0,0),
      ROW(0,0,0)
    ),
    ROW(
      ROW(0,0,0),
      ROW(0,0,0)
    ),
    ROW(false,false,false)
  ),
  try2 stage2_tries DEFAULT NULL,
  is_try2 boolean NOT NULL DEFAULT FALSE,
  shooter_sign varchar(255) DEFAULT NULL,
  scorer_sign varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE
);
