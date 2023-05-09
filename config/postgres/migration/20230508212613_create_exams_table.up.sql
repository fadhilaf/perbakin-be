CREATE TABLE IF NOT EXISTS exams (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  super_id uuid NOT NULL,
  name varchar(255) NOT NULL UNIQUE,
  location varchar(255) NOT NULL,
  organizer varchar(255) NOT NULL,
  begin date NOT NULL,
  finish date NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT super_id
    FOREIGN KEY (super_id) 
      REFERENCES supers (id)
);
