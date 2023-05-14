// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: shooters.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createShooter = `-- name: CreateShooter :one
INSERT INTO shooters (scorer_id, name, province, club)
VALUES ($1, $2, $3, $4)
RETURNING id, scorer_id, name, province, club, created_at, updated_at
`

type CreateShooterParams struct {
	ScorerID pgtype.UUID
	Name     string
	Province string
	Club     string
}

// membuat shooter baru (admin-super role)
func (q *Queries) CreateShooter(ctx context.Context, arg CreateShooterParams) (Shooter, error) {
	row := q.db.QueryRow(ctx, createShooter,
		arg.ScorerID,
		arg.Name,
		arg.Province,
		arg.Club,
	)
	var i Shooter
	err := row.Scan(
		&i.ID,
		&i.ScorerID,
		&i.Name,
		&i.Province,
		&i.Club,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteShooter = `-- name: DeleteShooter :exec
DELETE FROM shooters
WHERE id = $1
`

// untuk menghapus shooter berdasarkan id (admin-super role)
func (q *Queries) DeleteShooter(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteShooter, id)
	return err
}

const getAllShooters = `-- name: GetAllShooters :many
SELECT shooters.id, exams.name AS exam, shooters.name AS name, province, club
FROM shooters INNER JOIN scorers ON shooters.scorer_id = scorers.id INNER JOIN exams ON scorers.exam_id = exams.id
`

type GetAllShootersRow struct {
	ID       pgtype.UUID
	Exam     string
	Name     string
	Province string
	Club     string
}

// untuk mengambil seluruh shooter (admin-super role)
func (q *Queries) GetAllShooters(ctx context.Context) ([]GetAllShootersRow, error) {
	rows, err := q.db.Query(ctx, getAllShooters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllShootersRow
	for rows.Next() {
		var i GetAllShootersRow
		if err := rows.Scan(
			&i.ID,
			&i.Exam,
			&i.Name,
			&i.Province,
			&i.Club,
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

const getShooterByExamId = `-- name: GetShooterByExamId :many
SELECT shooters.id, users.name AS scorer, shooters.name AS name, province, club
FROM shooters INNER JOIN scorers ON shooters.scorer_id = scorers.id INNER JOIN users ON scorers.user_id = users.id
WHERE scorers.exam_id = $1
`

type GetShooterByExamIdRow struct {
	ID       pgtype.UUID
	Scorer   string
	Name     string
	Province string
	Club     string
}

// untuk mengambil shooter berdasarkan exam_id (admin-super role) TODO: tambah nilai results jg
func (q *Queries) GetShooterByExamId(ctx context.Context, examID pgtype.UUID) ([]GetShooterByExamIdRow, error) {
	rows, err := q.db.Query(ctx, getShooterByExamId, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetShooterByExamIdRow
	for rows.Next() {
		var i GetShooterByExamIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Scorer,
			&i.Name,
			&i.Province,
			&i.Club,
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

const getShooterById = `-- name: GetShooterById :one
SELECT id, scorer_id, name, province, club, created_at, updated_at
FROM shooters
WHERE id = $1
`

// untuk mengambil shooter berdasarkan id (admin-super role)
func (q *Queries) GetShooterById(ctx context.Context, id pgtype.UUID) (Shooter, error) {
	row := q.db.QueryRow(ctx, getShooterById, id)
	var i Shooter
	err := row.Scan(
		&i.ID,
		&i.ScorerID,
		&i.Name,
		&i.Province,
		&i.Club,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getShootersByScorerId = `-- name: GetShootersByScorerId :many
SELECT id, name, province, club
FROM shooters
WHERE scorer_id = $1
`

type GetShootersByScorerIdRow struct {
	ID       pgtype.UUID
	Name     string
	Province string
	Club     string
}

// untuk mengambil shooter berdasarkan scorer_id (admin-super role)
func (q *Queries) GetShootersByScorerId(ctx context.Context, scorerID pgtype.UUID) ([]GetShootersByScorerIdRow, error) {
	rows, err := q.db.Query(ctx, getShootersByScorerId, scorerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetShootersByScorerIdRow
	for rows.Next() {
		var i GetShootersByScorerIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Province,
			&i.Club,
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

const updateShooter = `-- name: UpdateShooter :one
UPDATE shooters 
SET scorer_id = $2, name = $3, province = $4, club = $5, updated_at = NOW()
WHERE id = $1
RETURNING id, scorer_id, name, province, club, created_at, updated_at
`

type UpdateShooterParams struct {
	ID       pgtype.UUID
	ScorerID pgtype.UUID
	Name     string
	Province string
	Club     string
}

// untuk mengupdate shooter berdasarkan id (admin-super role)
func (q *Queries) UpdateShooter(ctx context.Context, arg UpdateShooterParams) (Shooter, error) {
	row := q.db.QueryRow(ctx, updateShooter,
		arg.ID,
		arg.ScorerID,
		arg.Name,
		arg.Province,
		arg.Club,
	)
	var i Shooter
	err := row.Scan(
		&i.ID,
		&i.ScorerID,
		&i.Name,
		&i.Province,
		&i.Club,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
