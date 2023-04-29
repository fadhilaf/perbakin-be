CREATE TABLE IF NOT EXISTS shooters (
  id uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  name varchar(255) NOT NULL,
  profile_photo_path varchar(255) NOT NULL,
  province varchar(255) NOT NULL,
  club varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
);
