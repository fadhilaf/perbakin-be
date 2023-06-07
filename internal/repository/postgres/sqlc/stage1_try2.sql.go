// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage1_try2.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage1try2 = `-- name: CreateStage1try2 :one
WITH added_stage1_try2 AS (
  INSERT INTO stage13_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, no4, no5, no6, checkmarks
), updated_stage1_results AS (
  UPDATE stage1_results
  SET 
    try2_id = (SELECT id FROM added_stage1_try2),
    is_try2 = true,
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
), updated_stage1_try1 AS (
  UPDATE stage13_tries
  SET status = '7'
  WHERE id = (SELECT try1_id FROM updated_stage1_results)
)
SELECT 
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks
FROM added_stage1_try2
`

type CreateStage1try2Row struct {
	Status     Stage13Status
	No1        string
	No2        string
	No3        string
	No4        string
	No5        string
	No6        string
	Checkmarks string
}

func (q *Queries) CreateStage1try2(ctx context.Context, id pgtype.UUID) (CreateStage1try2Row, error) {
	row := q.db.QueryRow(ctx, createStage1try2, id)
	var i CreateStage1try2Row
	err := row.Scan(
		&i.Status,
		&i.No1,
		&i.No2,
		&i.No3,
		&i.No4,
		&i.No5,
		&i.No6,
		&i.Checkmarks,
	)
	return i, err
}

const getStage1try2Status = `-- name: GetStage1try2Status :one
SELECT 
  status
FROM stage1_results
INNER JOIN stage13_tries ON stage13_tries.id = stage1_results.try2_id
WHERE stage1_results.id = $1
`

// (all role)
func (q *Queries) GetStage1try2Status(ctx context.Context, id pgtype.UUID) (Stage13Status, error) {
	row := q.db.QueryRow(ctx, getStage1try2Status, id)
	var status Stage13Status
	err := row.Scan(&status)
	return status, err
}

const updateStage1try2 = `-- name: UpdateStage1try2 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1 
  RETURNING try2_id, updated_at
), updated_stage13_tries AS (
  UPDATE stage13_tries
  SET 
    status = $2, 
    no1 = $3, 
    no2 = $4, 
    no3 = $5, 
    no4 = $6,
    no5 = $7, 
    no6 = $8, 
    checkmarks = $9
  WHERE id = (SELECT try2_id FROM stage1_results)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    no4,
    no5,
    no6,
    checkmarks
)
SELECT 
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks,
  updated_at
FROM updated_stage1_results, updated_stage13_tries
`

type UpdateStage1try2Params struct {
	ID         pgtype.UUID
	Status     Stage13Status
	No1        string
	No2        string
	No3        string
	No4        string
	No5        string
	No6        string
	Checkmarks string
}

type UpdateStage1try2Row struct {
	Status     Stage13Status
	No1        string
	No2        string
	No3        string
	No4        string
	No5        string
	No6        string
	Checkmarks string
	UpdatedAt  pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage1try2(ctx context.Context, arg UpdateStage1try2Params) (UpdateStage1try2Row, error) {
	row := q.db.QueryRow(ctx, updateStage1try2,
		arg.ID,
		arg.Status,
		arg.No1,
		arg.No2,
		arg.No3,
		arg.No4,
		arg.No5,
		arg.No6,
		arg.Checkmarks,
	)
	var i UpdateStage1try2Row
	err := row.Scan(
		&i.Status,
		&i.No1,
		&i.No2,
		&i.No3,
		&i.No4,
		&i.No5,
		&i.No6,
		&i.Checkmarks,
		&i.UpdatedAt,
	)
	return i, err
}

const updateStage1try2Checkmarks = `-- name: UpdateStage1try2Checkmarks :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING checkmarks
`

type UpdateStage1try2CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage1try2Checkmarks(ctx context.Context, arg UpdateStage1try2CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage1try2FinishFailed = `-- name: UpdateStage1try2FinishFailed :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage1_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results)
`

type UpdateStage1try2FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage1try2FinishFailed(ctx context.Context, arg UpdateStage1try2FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage1try2FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage1try2FinishSuccess = `-- name: UpdateStage1try2FinishSuccess :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage1_results)
)
UPDATE results 
SET stage = '2', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results)
`

type UpdateStage1try2FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage1try2FinishSuccess(ctx context.Context, arg UpdateStage1try2FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage1try2FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage1try2NextNo = `-- name: UpdateStage1try2NextNo :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
  SET status = $2
WHERE id = (SELECT try2_id FROM stage1_results)
`

type UpdateStage1try2NextNoParams struct {
	ID     pgtype.UUID
	Status Stage13Status
}

// (scorer role)
func (q *Queries) UpdateStage1try2NextNo(ctx context.Context, arg UpdateStage1try2NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage1try2NextNo, arg.ID, arg.Status)
	return err
}

const updateStage1try2No1 = `-- name: UpdateStage1try2No1 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no1 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no1
`

type UpdateStage1try2No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage1try2No1(ctx context.Context, arg UpdateStage1try2No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage1try2No2 = `-- name: UpdateStage1try2No2 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no2 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no2
`

type UpdateStage1try2No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage1try2No2(ctx context.Context, arg UpdateStage1try2No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage1try2No3 = `-- name: UpdateStage1try2No3 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
) 
UPDATE stage13_tries 
SET no3 = $2 
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no3
`

type UpdateStage1try2No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage1try2No3(ctx context.Context, arg UpdateStage1try2No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}

const updateStage1try2No4 = `-- name: UpdateStage1try2No4 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no4 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no4
`

type UpdateStage1try2No4Params struct {
	ID  pgtype.UUID
	No4 string
}

// (scorer role)
func (q *Queries) UpdateStage1try2No4(ctx context.Context, arg UpdateStage1try2No4Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2No4, arg.ID, arg.No4)
	var no4 string
	err := row.Scan(&no4)
	return no4, err
}

const updateStage1try2No5 = `-- name: UpdateStage1try2No5 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no5 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no5
`

type UpdateStage1try2No5Params struct {
	ID  pgtype.UUID
	No5 string
}

// (scorer role)
func (q *Queries) UpdateStage1try2No5(ctx context.Context, arg UpdateStage1try2No5Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2No5, arg.ID, arg.No5)
	var no5 string
	err := row.Scan(&no5)
	return no5, err
}

const updateStage1try2No6 = `-- name: UpdateStage1try2No6 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no6 = $2
WHERE id = (SELECT try2_id FROM stage1_results)
RETURNING no6
`

type UpdateStage1try2No6Params struct {
	ID  pgtype.UUID
	No6 string
}

// (scorer role)
func (q *Queries) UpdateStage1try2No6(ctx context.Context, arg UpdateStage1try2No6Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try2No6, arg.ID, arg.No6)
	var no6 string
	err := row.Scan(&no6)
	return no6, err
}
