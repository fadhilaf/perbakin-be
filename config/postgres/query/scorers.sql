-- name: CreateScorer :exec
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($1, $2, $3)
  RETURNING id
)
INSERT INTO scorers (user_id)
SELECT id FROM added_user;

-- name: GetAllScorers :many
SELECT scorers.id, name, created_at, updated_at FROM scorers 
INNER JOIN users ON scorers.user_id = users.id;

-- name: GetScorerData :one
SELECT scorers.id, name, created_at, updated_at FROM scorers
INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.id = $1;

-- name: GetScorerByUserId :one
SELECT scorers.id, user_id, username, name FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
WHERE user_id = $1;

-- name: GetScorerByUsername :one
SELECT scorers.id, user_id, username, password, name FROM users
INNER JOIN scorers ON scorers.user_id = users.id
WHERE username = $1;

-- name: GetScorer :one
SELECT scorers.id, user_id, username, name, created_at, updated_at FROM scorers
INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.id = $1;

-- name: UpdateScorer :exec
UPDATE users 
SET username = $2, password = $3, name = $4, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM scorers 
  WHERE scorers.id = $1
);

-- name: UpdateScorerName :exec
UPDATE users 
SET name = $2, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM scorers 
  WHERE scorers.id = $1
);

-- name: UpdateScorerPassword :exec
UPDATE users 
SET password = $2, updated_at = NOW()
WHERE user_id = (
  SELECT user_id FROM scorers 
  WHERE scorers.id = $1
);

-- name: DeleteScorer :exec
DELETE FROM scorers WHERE user_id = $1;
