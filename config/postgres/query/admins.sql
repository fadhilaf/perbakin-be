-- name: CreateAdmin :execresult
INSERT INTO "users" (
  "username", "password", "name"
)
VALUES (
  $1, $2, $3
)
RETURNING "id" INTO "admins" (user_id);

-- name: GetAdmin :one
SELECT "id", "user_id" FROM "users"
WHERE "user_id" = $1;
