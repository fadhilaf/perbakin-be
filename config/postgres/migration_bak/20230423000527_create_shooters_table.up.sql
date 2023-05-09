CREATE TABLE IF NOT EXISTS shooters (
  id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
  name varchar(255) NOT NULL,
  image_path varchar(255) NOT NULL,
  province varchar(255) NOT NULL,
  club varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
