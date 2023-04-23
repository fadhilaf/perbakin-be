-- migrate create -ext sql create_users_table

DROP TABLE IF EXISTS "scorers";
DROP TABLE IF EXISTS "admins";
DROP TABLE IF EXISTS "supers";

DROP TABLE IF EXISTS "users";

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = "user_role") THEN
        DROP TYPE "user_role";
    END IF;
END $$;
