DO $$
  BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = "stages") THEN
      CREATE TYPE "stages" AS ENUM ("0", "1", "2", "3", "4", "5", "6");
    END IF;
  END
$$;

CREATE TABLE IF NOT EXISTS "results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "shooter_test_id" uuid NOT NULL,
  "status" boolean NOT NULL default false,
  "stage" "stages" NOT NULL DEFAULT "0",
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "shooter_test_id"
    FOREIGN KEY ("shooter_test_id") 
      REFERENCES "shooter_tests" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage0_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage1_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage2_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage3_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage4_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage5_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)

CREATE TABLE IF NOT EXISTS "stage6_results" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "result_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "result_id"
    FOREIGN KEY ("result_id") 
      REFERENCES "results" ("id"),
)
