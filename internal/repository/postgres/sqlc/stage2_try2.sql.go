// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage2_try2.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage2try2 = `-- name: CreateStage2try2 :one
WITH added_stage2_try2 AS (
  INSERT INTO stage2_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), updated_stage2_results AS (
  UPDATE stage2_results
  SET 
    try2_id = (SELECT id FROM added_stage2_try2),
    is_try2 = true,
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id, is_try2
), updated_stage2_try1 AS (
  UPDATE stage2_tries
  SET status = '4'
  WHERE id = (SELECT try1_id FROM updated_stage2_results)
)
SELECT 
  is_try2,
  status,
  no1,
  no2,
  no3,
  checkmarks
FROM added_stage2_try2, updated_stage2_results
`

type CreateStage2try2Row struct {
	IsTry2     bool
	Status     Stage246Status
	No1        string
	No2        string
	No3        string
	Checkmarks string
}

func (q *Queries) CreateStage2try2(ctx context.Context, id pgtype.UUID) (CreateStage2try2Row, error) {
	row := q.db.QueryRow(ctx, createStage2try2, id)
	var i CreateStage2try2Row
	err := row.Scan(
		&i.IsTry2,
		&i.Status,
		&i.No1,
		&i.No2,
		&i.No3,
		&i.Checkmarks,
	)
	return i, err
}

const deleteStage2try2 = `-- name: DeleteStage2try2 :exec
WITH deleted_stage2_try2 AS (
  DELETE FROM stage2_tries
  WHERE stage2_tries.id = (SELECT try2_id FROM stage2_results WHERE stage2_results.id = $1)
), updated_stage2_results AS (
  UPDATE stage2_results
  SET
    try2_id = NULL,
    is_try2 = false,
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET status = '6'
WHERE stage2_tries.id = (SELECT try1_id FROM updated_stage2_results)
`

// (admin-super role)
func (q *Queries) DeleteStage2try2(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteStage2try2, id)
	return err
}

const getStage2try2Status = `-- name: GetStage2try2Status :one
SELECT 
  status
FROM stage2_results
INNER JOIN stage2_tries ON stage2_tries.id = stage2_results.try2_id
WHERE stage2_results.id = $1
`

// (all role)
func (q *Queries) GetStage2try2Status(ctx context.Context, id pgtype.UUID) (Stage246Status, error) {
	row := q.db.QueryRow(ctx, getStage2try2Status, id)
	var status Stage246Status
	err := row.Scan(&status)
	return status, err
}

const updateStage2try2 = `-- name: UpdateStage2try2 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage2_try1 AS (
  UPDATE stage2_tries
  SET 
    status = $2,
    no1 = $3,
    no2 = $4,
    no3 = $5,
    checkmarks = $6
  WHERE id = (SELECT try1_id FROM stage2_results)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
), updated_stage2_try2 AS (
  UPDATE stage2_tries
  SET 
    status = $7,
    no1 = $8,
    no2 = $9,
    no3 = $10,
    checkmarks = $11
  WHERE id = (SELECT try2_id FROM stage2_results WHERE try2_id IS NOT NULL)
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
)
SELECT 
  updated_stage2_try1.status AS try1_status,
  updated_stage2_try1.no1 AS try1_no1,
  updated_stage2_try1.no2 AS try1_no2,
  updated_stage2_try1.no3 AS try1_no3,
  updated_stage2_try1.checkmarks AS try1_checkmarks,
  updated_stage2_try2.status AS try2_status,
  updated_stage2_try2.no1 AS try2_no1,
  updated_stage2_try2.no2 AS try2_no2,
  updated_stage2_try2.no3 AS try2_no3,
  updated_stage2_try2.checkmarks AS try2_checkmarks,
  updated_at
FROM updated_stage2_results, updated_stage2_try1, updated_stage2_try2
`

type UpdateStage2try2Params struct {
	ID             pgtype.UUID
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
	Try2Status     Stage246Status
	Try2No1        string
	Try2No2        string
	Try2No3        string
	Try2Checkmarks string
}

type UpdateStage2try2Row struct {
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
	Try2Status     Stage246Status
	Try2No1        string
	Try2No2        string
	Try2No3        string
	Try2Checkmarks string
	UpdatedAt      pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage2try2(ctx context.Context, arg UpdateStage2try2Params) (UpdateStage2try2Row, error) {
	row := q.db.QueryRow(ctx, updateStage2try2,
		arg.ID,
		arg.Try1Status,
		arg.Try1No1,
		arg.Try1No2,
		arg.Try1No3,
		arg.Try1Checkmarks,
		arg.Try2Status,
		arg.Try2No1,
		arg.Try2No2,
		arg.Try2No3,
		arg.Try2Checkmarks,
	)
	var i UpdateStage2try2Row
	err := row.Scan(
		&i.Try1Status,
		&i.Try1No1,
		&i.Try1No2,
		&i.Try1No3,
		&i.Try1Checkmarks,
		&i.Try2Status,
		&i.Try2No1,
		&i.Try2No2,
		&i.Try2No3,
		&i.Try2Checkmarks,
		&i.UpdatedAt,
	)
	return i, err
}

const updateStage2try2Checkmarks = `-- name: UpdateStage2try2Checkmarks :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try2_id FROM stage2_results)
RETURNING checkmarks
`

type UpdateStage2try2CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage2try2Checkmarks(ctx context.Context, arg UpdateStage2try2CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try2Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage2try2FinishFailed = `-- name: UpdateStage2try2FinishFailed :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try2_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '4'
  WHERE id = (SELECT try2_id FROM updated_stage2_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage2_results)
`

type UpdateStage2try2FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage2try2FinishFailed(ctx context.Context, arg UpdateStage2try2FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage2try2FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage2try2FinishSuccess = `-- name: UpdateStage2try2FinishSuccess :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try2_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '4'
  WHERE id = (SELECT try2_id FROM updated_stage2_results)
)
UPDATE results 
SET stage = '3', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage2_results)
`

type UpdateStage2try2FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage2try2FinishSuccess(ctx context.Context, arg UpdateStage2try2FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage2try2FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage2try2NextNo = `-- name: UpdateStage2try2NextNo :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
  SET status = $2
WHERE id = (SELECT try2_id FROM stage2_results)
`

type UpdateStage2try2NextNoParams struct {
	ID     pgtype.UUID
	Status Stage246Status
}

// (scorer role)
func (q *Queries) UpdateStage2try2NextNo(ctx context.Context, arg UpdateStage2try2NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage2try2NextNo, arg.ID, arg.Status)
	return err
}

const updateStage2try2No1 = `-- name: UpdateStage2try2No1 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
SET no1 = $2
WHERE id = (SELECT try2_id FROM stage2_results)
RETURNING no1
`

type UpdateStage2try2No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage2try2No1(ctx context.Context, arg UpdateStage2try2No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try2No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage2try2No2 = `-- name: UpdateStage2try2No2 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
)
UPDATE stage2_tries
SET no2 = $2
WHERE id = (SELECT try2_id FROM stage2_results)
RETURNING no2
`

type UpdateStage2try2No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage2try2No2(ctx context.Context, arg UpdateStage2try2No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try2No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage2try2No3 = `-- name: UpdateStage2try2No3 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try2_id
) 
UPDATE stage2_tries 
SET no3 = $2 
WHERE id = (SELECT try2_id FROM stage2_results)
RETURNING no3
`

type UpdateStage2try2No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage2try2No3(ctx context.Context, arg UpdateStage2try2No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try2No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}
