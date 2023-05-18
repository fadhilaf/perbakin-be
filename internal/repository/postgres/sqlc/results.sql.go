// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: results.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createResult = `-- name: CreateResult :one
INSERT INTO results (shooter_id)
VALUES ($1)
RETURNING id, shooter_id, failed, stage, created_at, updated_at
`

func (q *Queries) CreateResult(ctx context.Context, shooterID pgtype.UUID) (Result, error) {
	row := q.db.QueryRow(ctx, createResult, shooterID)
	var i Result
	err := row.Scan(
		&i.ID,
		&i.ShooterID,
		&i.Failed,
		&i.Stage,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getResultByShooterId = `-- name: GetResultByShooterId :one
SELECT id, shooter_id, failed, stage, created_at, updated_at
FROM results
WHERE shooter_id = $1
`

func (q *Queries) GetResultByShooterId(ctx context.Context, shooterID pgtype.UUID) (Result, error) {
	row := q.db.QueryRow(ctx, getResultByShooterId, shooterID)
	var i Result
	err := row.Scan(
		&i.ID,
		&i.ShooterID,
		&i.Failed,
		&i.Stage,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
