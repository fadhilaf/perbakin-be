-- migrate create -ext sql create_users_table

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(255) NOT NULL,
  "username" varchar(255) NOT NULL UNIQUE,
  "email" varchar(255),
  "password" varchar(255) NOT NULL,
  "email_verified" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) 

CREATE TABLE IF NOT EXISTS "supers" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  CONSTRAINT "user_id"
    FOREIGN KEY ("user_id") 
      REFERENCES "users" ("id")
)

CREATE TABLE IF NOT EXISTS "admins" ( 
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  CONSTRAINT "user_id"
    FOREIGN KEY ("user_id") 
      REFERENCES "users" ("id")
)

CREATE TABLE IF NOT EXISTS "scorers" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  CONSTRAINT "user_id"
    FOREIGN KEY ("user_id") 
      REFERENCES "users" ("id")
)
