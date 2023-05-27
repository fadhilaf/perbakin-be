CREATE TYPE stage0_series AS (
  score0 integer,
  score1 integer,
  score2 integer,
  score3 integer,
  score4 integer,
  score5 integer,
  score6 integer,
  score7 integer,
  score8 integer,
  score9 integer,
  score10 integer
);

CREATE TYPE stage0_checkmarks AS (
  check1 boolean,
  check2 boolean,
  check3 boolean,
  check4 boolean,
  check5 boolean
);

CREATE TYPE stage0_status AS ENUM ('1', '2', '3', '4', '5', '6');

CREATE TABLE IF NOT EXISTS stage0_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  status stage0_status NOT NULL DEFAULT '1',
  series1 stage0_series NOT NULL DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series2 stage0_series NOT NULL DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series3 stage0_series NOT NULL DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series4 stage0_series NOT NULL DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series5 stage0_series NOT NULL DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  checkmarks stage0_checkmarks NOT NULL DEFAULT (ROW(false,false,false,false,false)),
  shooter_sign varchar(255) DEFAULT NULL,
  scorer_sign varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE
);
