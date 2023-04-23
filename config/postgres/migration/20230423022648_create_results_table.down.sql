DROP TABLE IF EXISTS "stage6_results";
DROP TABLE IF EXISTS "stage5_results";
DROP TABLE IF EXISTS "stage4_results";
DROP TABLE IF EXISTS "stage3_results";
DROP TABLE IF EXISTS "stage2_results";
DROP TABLE IF EXISTS "stage1_results";
DROP TABLE IF EXISTS "stage0_results";

DROP TABLE IF EXISTS "results";

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = "stages") THEN
        DROP TYPE "stages";
    END IF;
END $$;
