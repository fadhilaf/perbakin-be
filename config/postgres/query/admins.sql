-- name: CreateAdmin :execresult
INSERT INTO "users" (
  "username", "email", "password", "name"
)
VALUES (
  $1, $2, $3, $4
)
RETURNING "id" INTO "admins" (user_id);

-- name: GetAdmin :one
SELECT "id", "user_id" FROM "users"
WHERE "user_id" = $1;
