// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: scorers.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createScorer = `-- name: CreateScorer :execresult
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($1, $2, $3)
  RETURNING id
)
INSERT INTO scorers (user_id)
SELECT id FROM added_user
`

type CreateScorerParams struct {
	Username string
	Password string
	Name     string
}

func (q *Queries) CreateScorer(ctx context.Context, arg CreateScorerParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createScorer, arg.Username, arg.Password, arg.Name)
}

const deleteScorer = `-- name: DeleteScorer :exec
DELETE FROM scorers WHERE user_id = $1
`

func (q *Queries) DeleteScorer(ctx context.Context, userID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteScorer, userID)
	return err
}

const getScorer = `-- name: GetScorer :one
SELECT scorers.id, user_id FROM scorers
WHERE user_id = $1
`

func (q *Queries) GetScorer(ctx context.Context, userID pgtype.UUID) (Scorer, error) {
	row := q.db.QueryRow(ctx, getScorer, userID)
	var i Scorer
	err := row.Scan(&i.ID, &i.UserID)
	return i, err
}

const getScorerByUsername = `-- name: GetScorerByUsername :one
SELECT scorers.id, user_id, username, password, name FROM users
INNER JOIN scorers ON scorers.user_id = users.id
WHERE username = $1
`

type GetScorerByUsernameRow struct {
	ID       pgtype.UUID
	UserID   pgtype.UUID
	Username string
	Password string
	Name     string
}

func (q *Queries) GetScorerByUsername(ctx context.Context, username string) (GetScorerByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getScorerByUsername, username)
	var i GetScorerByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Name,
	)
	return i, err
}

const getScorers = `-- name: GetScorers :many
SELECT scorers.id, user_id FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
`

func (q *Queries) GetScorers(ctx context.Context) ([]Scorer, error) {
	rows, err := q.db.Query(ctx, getScorers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Scorer
	for rows.Next() {
		var i Scorer
		if err := rows.Scan(&i.ID, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateScorer = `-- name: UpdateScorer :exec
UPDATE users SET username = $2, name = $3 WHERE id = $1
`

type UpdateScorerParams struct {
	ID       pgtype.UUID
	Username string
	Name     string
}

func (q *Queries) UpdateScorer(ctx context.Context, arg UpdateScorerParams) error {
	_, err := q.db.Exec(ctx, updateScorer, arg.ID, arg.Username, arg.Name)
	return err
}
