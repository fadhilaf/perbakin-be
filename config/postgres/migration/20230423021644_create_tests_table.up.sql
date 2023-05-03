CREATE TABLE IF NOT EXISTS shooting_tests (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  super_id uuid NOT NULL,
  name varchar(255) NOT NULL,
  location varchar(255) NOT NULL,
  begin timestamp NOT NULL,
  finish timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT super_id
    FOREIGN KEY (super_id) 
      REFERENCES supers (id)
);

CREATE TABLE IF NOT EXISTS admin_tests (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  admin_id uuid NOT NULL,
  test_id uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT admin_id
    FOREIGN KEY (admin_id) 
      REFERENCES admins (id),
  CONSTRAINT test_id
    FOREIGN KEY (test_id) 
      REFERENCES shooting_tests (id)
);

CREATE TABLE IF NOT EXISTS scorer_tests (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  scorer_id uuid NOT NULL,
  test_id uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT scorer_id
    FOREIGN KEY (scorer_id) 
      REFERENCES scorers (id),
  CONSTRAINT test_id
    FOREIGN KEY (test_id) 
      REFERENCES shooting_tests (id)
);

CREATE TABLE IF NOT EXISTS shooter_tests (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  shooter_id uuid NOT NULL,
  scorer_test_id uuid NOT NULL,
  created_by uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT shooter_id
    FOREIGN KEY (shooter_id) 
      REFERENCES shooters (id),
  CONSTRAINT scorer_test_id
    FOREIGN KEY (scorer_test_id) 
      REFERENCES scorer_tests (id),
  CONSTRAINT created_by
    FOREIGN KEY (created_by)
      REFERENCES admin_tests (id)
);
