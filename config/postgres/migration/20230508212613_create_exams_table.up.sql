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

CREATE TABLE IF NOT EXISTS admins ( 
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  user_id uuid NOT NULL,
  exam_id uuid NOT NULL,
  CONSTRAINT user_id
    FOREIGN KEY (user_id) 
      REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT exam_id
    FOREIGN KEY (exam_id) 
      REFERENCES exams (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS scorers (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  user_id uuid NOT NULL,
  exam_id uuid NOT NULL,
  CONSTRAINT user_id
    FOREIGN KEY (user_id) 
      REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT exam_id
    FOREIGN KEY (exam_id) 
      REFERENCES exams (id) ON DELETE CASCADE
);
