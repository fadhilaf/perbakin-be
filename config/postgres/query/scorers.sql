-- name: CreateScorer :execresult
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($1, $2, $3)
  RETURNING id
)
INSERT INTO scorers (user_id)
SELECT id FROM added_user;

-- name: GetScorers :many
SELECT scorers.id, user_id FROM scorers 
INNER JOIN users ON scorers.user_id = users.id;

-- name: GetScorer :one
SELECT scorers.id, user_id FROM scorers
WHERE user_id = $1;

-- name: GetScorerByUsername :one
SELECT scorers.id, user_id, username, name FROM users
INNER JOIN scorers ON scorers.user_id = users.id
WHERE username = $1;

-- name: DeleteScorer :exec
DELETE FROM scorers WHERE user_id = $1;
