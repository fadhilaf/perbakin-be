CREATE TYPE stage123456_durations AS (
  minute integer,
  second integer,
  millisecond integer
);

CREATE TYPE stage123456_scores AS (
  score_a integer,
  score_c integer,
  score_d integer
);

CREATE TYPE stage123_numbers AS (
  scores stage123456_scores,
  duration stage123456_durations
);

CREATE TYPE stage13_checkmarks AS (
  check1 boolean,
  check2 boolean,
  check3 boolean,
  check4 boolean,
  check5 boolean,
  check6 boolean
);

CREATE TYPE stage13_status AS ENUM ('1', '2', '3', '4', '5', '6', '7');

CREATE TYPE stage13_tries AS (
  status stage13_status,
  no1 stage123_numbers,
  no2 stage123_numbers,
  no3 stage123_numbers,
  no4 stage123_numbers,
  no5 stage123_numbers,
  no6 stage123_numbers,
  checkmarks stage13_checkmarks
);

CREATE TABLE IF NOT EXISTS stage1_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  try1 stage13_tries NOT NULL DEFAULT ROW( 
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
    ROW(false,false,false,false,false,false)
  ),
  try2 stage13_tries DEFAULT NULL,
  is_try2 boolean NOT NULL DEFAULT FALSE,
  shooter_sign varchar(255) DEFAULT NULL,
  scorer_sign varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE
);
