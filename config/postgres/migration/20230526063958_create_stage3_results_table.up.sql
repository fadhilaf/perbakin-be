CREATE TABLE IF NOT EXISTS stage3_results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  result_id uuid NOT NULL UNIQUE,
  try1 stage13_tries NOT NULL DEFAULT ROW( 
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
