DO $$
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'stages') THEN
      CREATE TYPE stages AS ENUM ('0', '1', '2', '3', '4', '5', '6', '7');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'stage0_series') THEN
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
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'stage0_status') THEN
      CREATE TYPE stage0_status AS ENUM ('1', '2', '3', '4', '5', '6');
    END IF;
  END
$$;

CREATE TABLE IF NOT EXISTS results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  shooter_id uuid NOT NULL UNIQUE,
  failed boolean NOT NULL DEFAULT false,
  stage stages DEFAULT NOT NULL DEFAULT '0',
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT shooter_id
    FOREIGN KEY (shooter_id) 
      REFERENCES shooters (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS stage0_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  status stage0_status NOT NULL DEFAULT '1',
  series1 stage0_series DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series2 stage0_series DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series3 stage0_series DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series4 stage0_series DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  series5 stage0_series DEFAULT (ROW(0,0,0,0,0,0,0,0,0,0,0)),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT result_id
    FOREIGN KEY (result_id) 
      REFERENCES results (id) ON DELETE CASCADE
);
