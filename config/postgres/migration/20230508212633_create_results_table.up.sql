CREATE TYPE stages AS ENUM ('0', '1', '2', '3', '4', '5', '6', '7');

CREATE TABLE IF NOT EXISTS results (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  shooter_id uuid NOT NULL UNIQUE,
  failed boolean NOT NULL DEFAULT false,
  stage stages NOT NULL DEFAULT '0',
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT shooter_id
    FOREIGN KEY (shooter_id) 
      REFERENCES shooters (id) ON DELETE CASCADE
);
