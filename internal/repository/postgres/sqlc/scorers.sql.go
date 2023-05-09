// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: scorers.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createScorer = `-- name: CreateScorer :one
WITH added_user AS (
  INSERT INTO users (username, password, name)
  VALUES ($2, $3, $4)
  RETURNING id
), added_scorer AS (
  INSERT INTO scorers (user_id, exam_id)
  SELECT id, $1 FROM added_user
  RETURNING id
)
SELECT scorers.id, user_id, exam_id, username, name, created_at, updated_at FROM scorers
INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.id = (
  SELECT id FROM added_scorer
)
`

type CreateScorerParams struct {
	ExamID   pgtype.UUID
	Username string
	Password string
	Name     string
}

type CreateScorerRow struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	ExamID    pgtype.UUID
	Username  string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// untuk ngebuat scorer (admin-super role) TODO: return sebanyak get scorer by id
func (q *Queries) CreateScorer(ctx context.Context, arg CreateScorerParams) (CreateScorerRow, error) {
	row := q.db.QueryRow(ctx, createScorer,
		arg.ExamID,
		arg.Username,
		arg.Password,
		arg.Name,
	)
	var i CreateScorerRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ExamID,
		&i.Username,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllScorers = `-- name: GetAllScorers :many
SELECT scorers.id, exams.name AS exam, users.name AS name, users.created_at, users.updated_at FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
INNER JOIN exams ON scorers.exam_id = exams.id
`

type GetAllScorersRow struct {
	ID        pgtype.UUID
	Exam      string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// untuk ngambil data display seluruh scorer (all role)
func (q *Queries) GetAllScorers(ctx context.Context) ([]GetAllScorersRow, error) {
	rows, err := q.db.Query(ctx, getAllScorers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllScorersRow
	for rows.Next() {
		var i GetAllScorersRow
		if err := rows.Scan(
			&i.ID,
			&i.Exam,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getScorerById = `-- name: GetScorerById :one
SELECT scorers.id, user_id, exam_id, username, name, created_at, updated_at FROM scorers
INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.id = $1
`

type GetScorerByIdRow struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	ExamID    pgtype.UUID
	Username  string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// untuk ngambil data akun scorer berdasarkan id (admin-super role)
func (q *Queries) GetScorerById(ctx context.Context, id pgtype.UUID) (GetScorerByIdRow, error) {
	row := q.db.QueryRow(ctx, getScorerById, id)
	var i GetScorerByIdRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ExamID,
		&i.Username,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getScorerByUsername = `-- name: GetScorerByUsername :one
SELECT scorers.id, user_id, exam_id, username, password, name, created_at, updated_at FROM users
INNER JOIN scorers ON scorers.user_id = users.id
WHERE username = $1
`

type GetScorerByUsernameRow struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	ExamID    pgtype.UUID
	Username  string
	Password  string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// untuk ngambil data display scorer berdasarkan username (scorer role)
func (q *Queries) GetScorerByUsername(ctx context.Context, username string) (GetScorerByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getScorerByUsername, username)
	var i GetScorerByUsernameRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ExamID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getScorerRelationByUserId = `-- name: GetScorerRelationByUserId :one
SELECT scorers.id, user_id, exam_id FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
WHERE user_id = $1
`

// untuk ngambil data relasi scorer berdasarkan user id (all role)
func (q *Queries) GetScorerRelationByUserId(ctx context.Context, userID pgtype.UUID) (Scorer, error) {
	row := q.db.QueryRow(ctx, getScorerRelationByUserId, userID)
	var i Scorer
	err := row.Scan(&i.ID, &i.UserID, &i.ExamID)
	return i, err
}

const getScorersByExamId = `-- name: GetScorersByExamId :many
SELECT scorers.id, user_id, exam_id, username, name, created_at, updated_at FROM scorers 
INNER JOIN users ON scorers.user_id = users.id
WHERE exam_id = $1
`

type GetScorersByExamIdRow struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	ExamID    pgtype.UUID
	Username  string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// untuk ngambil data akun seluruh scorer dalam satu exam (admin-super role)
func (q *Queries) GetScorersByExamId(ctx context.Context, examID pgtype.UUID) ([]GetScorersByExamIdRow, error) {
	rows, err := q.db.Query(ctx, getScorersByExamId, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetScorersByExamIdRow
	for rows.Next() {
		var i GetScorersByExamIdRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ExamID,
			&i.Username,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateScorer = `-- name: UpdateScorer :one
WITH updated_user AS (
  UPDATE users 
  SET username = $2, password = $3, name = $4, updated_at = NOW() 
  WHERE users.id = (
    SELECT user_id FROM scorers 
    WHERE scorers.id = $1
  )
  RETURNING id
)
SELECT scorers.id, user_id, exam_id, username, name, created_at, updated_at FROM scorers
INNER JOIN users ON scorers.user_id = users.id
WHERE user_id = (
  SELECT id FROM updated_user
)
`

type UpdateScorerParams struct {
	ID       pgtype.UUID
	Username string
	Password string
	Name     string
}

type UpdateScorerRow struct {
	ID        pgtype.UUID
	UserID    pgtype.UUID
	ExamID    pgtype.UUID
	Username  string
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

// untuk update data akun admin (super role) TODO: return sebanyak get admin by id
func (q *Queries) UpdateScorer(ctx context.Context, arg UpdateScorerParams) (UpdateScorerRow, error) {
	row := q.db.QueryRow(ctx, updateScorer,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Name,
	)
	var i UpdateScorerRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ExamID,
		&i.Username,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateScorerName = `-- name: UpdateScorerName :one
UPDATE users 
SET name = $2, updated_at = NOW() 
WHERE user_id = (
  SELECT user_id FROM scorers 
  WHERE scorers.id = $1
)
RETURNING id
`

type UpdateScorerNameParams struct {
	ID   pgtype.UUID
	Name string
}

// low prio
func (q *Queries) UpdateScorerName(ctx context.Context, arg UpdateScorerNameParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, updateScorerName, arg.ID, arg.Name)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const updateScorerPassword = `-- name: UpdateScorerPassword :one
UPDATE users 
SET password = $2, updated_at = NOW()
WHERE user_id = (
  SELECT user_id FROM scorers 
  WHERE scorers.id = $1
)
RETURNING id
`

type UpdateScorerPasswordParams struct {
	ID       pgtype.UUID
	Password string
}

// low prio
func (q *Queries) UpdateScorerPassword(ctx context.Context, arg UpdateScorerPasswordParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, updateScorerPassword, arg.ID, arg.Password)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}
