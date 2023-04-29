-- migrate create -ext sql create_users_table

-- bisa jadi tidak perlu email, biar lebih simpel

CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  username varchar(255) NOT NULL UNIQUE,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS supers (
  id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  user_id uuid NOT NULL,
  CONSTRAINT user_id
    FOREIGN KEY (user_id) 
      REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS admins ( 
  id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  user_id uuid NOT NULL,
  CONSTRAINT user_id
    FOREIGN KEY (user_id) 
      REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS scorers (
  id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  user_id uuid NOT NULL,
  CONSTRAINT user_id
    FOREIGN KEY (user_id) 
      REFERENCES users (id)
);
