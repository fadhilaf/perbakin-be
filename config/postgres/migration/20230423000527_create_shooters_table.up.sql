CREATE TABLE IF NOT EXISTS "shooters" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(255) NOT NULL,
  "profile_photo_path" varchar(255) NOT NULL,
  "province" varchar(255) NOT NULL,
  "club" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
)
