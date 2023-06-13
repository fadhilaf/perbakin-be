CREATE TABLE IF NOT EXISTS stage3_results (
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
      REFERENCES stage13_tries (id) ON DELETE CASCADE,
  CONSTRAINT try2_id 
    FOREIGN KEY (try2_id) 
      REFERENCES stage13_tries (id) ON DELETE SET NULL
);
