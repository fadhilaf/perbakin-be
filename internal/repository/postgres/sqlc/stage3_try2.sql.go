// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage3_try2.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage3try2 = `-- name: CreateStage3try2 :one
WITH added_stage3_try2 AS (
  INSERT INTO stage13_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, no4, no5, no6, checkmarks
), updated_stage3_results AS (
  UPDATE stage3_results
  SET 
    try2_id = (SELECT id FROM added_stage3_try2),
    is_try2 = true,
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try1_id, is_try2
), updated_stage3_try1 AS (
  UPDATE stage13_tries
  SET status = '7'
  WHERE id = updated_stage3_results.try1_id
)
SELECT 
  is_try2,
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks
FROM added_stage3_try2, updated_stage3_results
`

type CreateStage3try2Row struct {
	IsTry2     bool
	Status     Stage13Status
	No1        string
	No2        string
	No3        string
	No4        string
	No5        string
	No6        string
	Checkmarks string
}

func (q *Queries) CreateStage3try2(ctx context.Context, id pgtype.UUID) (CreateStage3try2Row, error) {
	row := q.db.QueryRow(ctx, createStage3try2, id)
	var i CreateStage3try2Row
	err := row.Scan(
		&i.IsTry2,
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

const deleteStage3try2 = `-- name: DeleteStage3try2 :exec
WITH deleted_stage3_try2 AS (
  DELETE FROM stage13_tries
  WHERE stage13_tries.id = (SELECT try2_id FROM stage3_results WHERE stage3_results.id = $1)
), updated_stage3_results AS (
  UPDATE stage3_results
  SET
    try2_id = NULL,
    is_try2 = false,
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET status = '6'
WHERE stage13_tries.id = updated_stage3_results.try1_id
`

// (admin-super role)
func (q *Queries) DeleteStage3try2(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteStage3try2, id)
	return err
}

const getStage3try2Status = `-- name: GetStage3try2Status :one
SELECT 
  status
FROM stage3_results
INNER JOIN stage13_tries ON stage13_tries.id = stage3_results.try2_id
WHERE stage3_results.id = $1
`

// (all role)
func (q *Queries) GetStage3try2Status(ctx context.Context, id pgtype.UUID) (Stage13Status, error) {
	row := q.db.QueryRow(ctx, getStage3try2Status, id)
	var status Stage13Status
	err := row.Scan(&status)
	return status, err
}

const updateStage3try2 = `-- name: UpdateStage3try2 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage3_try1 AS (
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
  WHERE id = updated_stage3_results.try1_id
  RETURNING 
    status,
    no1,
    no2,
    no3,
    no4,
    no5,
    no6,
    checkmarks
), updated_stage3_try2 AS (
  UPDATE stage13_tries
  SET 
    status = $10,
    no1 = $11,
    no2 = $12,
    no3 = $13,
    no4 = $14,
    no5 = $15,
    no6 = $16,
    checkmarks = $17
  WHERE id = (SELECT try2_id FROM updated_stage3_results WHERE try2_id IS NOT NULL)
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
  updated_stage3_try1.status AS try1_status,
  updated_stage3_try1.no1 AS try1_no1,
  updated_stage3_try1.no2 AS try1_no2,
  updated_stage3_try1.no3 AS try1_no3,
  updated_stage3_try1.no4 AS try1_no4,
  updated_stage3_try1.no5 AS try1_no5,
  updated_stage3_try1.no6 AS try1_no6,
  updated_stage3_try1.checkmarks AS try1_checkmarks,
  updated_stage3_try2.status AS try2_status,
  updated_stage3_try2.no1 AS try2_no1,
  updated_stage3_try2.no2 AS try2_no2,
  updated_stage3_try2.no3 AS try2_no3,
  updated_stage3_try2.no4 AS try2_no4,
  updated_stage3_try2.no5 AS try2_no5,
  updated_stage3_try2.no6 AS try2_no6,
  updated_stage3_try2.checkmarks AS try2_checkmarks,
  updated_at
FROM updated_stage3_results, updated_stage3_try1, updated_stage3_try2
`

type UpdateStage3try2Params struct {
	ID             pgtype.UUID
	Try1Status     Stage13Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1No4        string
	Try1No5        string
	Try1No6        string
	Try1Checkmarks string
	Try2Status     Stage13Status
	Try2No1        string
	Try2No2        string
	Try2No3        string
	Try2No4        string
	Try2No5        string
	Try2No6        string
	Try2Checkmarks string
}

type UpdateStage3try2Row struct {
	Try1Status     Stage13Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1No4        string
	Try1No5        string
	Try1No6        string
	Try1Checkmarks string
	Try2Status     Stage13Status
	Try2No1        string
	Try2No2        string
	Try2No3        string
	Try2No4        string
	Try2No5        string
	Try2No6        string
	Try2Checkmarks string
	UpdatedAt      pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage3try2(ctx context.Context, arg UpdateStage3try2Params) (UpdateStage3try2Row, error) {
	row := q.db.QueryRow(ctx, updateStage3try2,
		arg.ID,
		arg.Try1Status,
		arg.Try1No1,
		arg.Try1No2,
		arg.Try1No3,
		arg.Try1No4,
		arg.Try1No5,
		arg.Try1No6,
		arg.Try1Checkmarks,
		arg.Try2Status,
		arg.Try2No1,
		arg.Try2No2,
		arg.Try2No3,
		arg.Try2No4,
		arg.Try2No5,
		arg.Try2No6,
		arg.Try2Checkmarks,
	)
	var i UpdateStage3try2Row
	err := row.Scan(
		&i.Try1Status,
		&i.Try1No1,
		&i.Try1No2,
		&i.Try1No3,
		&i.Try1No4,
		&i.Try1No5,
		&i.Try1No6,
		&i.Try1Checkmarks,
		&i.Try2Status,
		&i.Try2No1,
		&i.Try2No2,
		&i.Try2No3,
		&i.Try2No4,
		&i.Try2No5,
		&i.Try2No6,
		&i.Try2Checkmarks,
		&i.UpdatedAt,
	)
	return i, err
}

const updateStage3try2Checkmarks = `-- name: UpdateStage3try2Checkmarks :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET 
  checkmarks = $2
WHERE id = updated_stage3_results.try2_id
RETURNING checkmarks
`

type UpdateStage3try2CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage3try2Checkmarks(ctx context.Context, arg UpdateStage3try2CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage3try2FinishFailed = `-- name: UpdateStage3try2FinishFailed :exec
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage3_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage3_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage3_results)
`

type UpdateStage3try2FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage3try2FinishFailed(ctx context.Context, arg UpdateStage3try2FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage3try2FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage3try2FinishSuccess = `-- name: UpdateStage3try2FinishSuccess :exec
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage3_results.id = $1
  RETURNING result_id, try2_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try2_id FROM updated_stage3_results)
)
UPDATE results 
SET stage = '4', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage3_results)
`

type UpdateStage3try2FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage3try2FinishSuccess(ctx context.Context, arg UpdateStage3try2FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage3try2FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage3try2NextNo = `-- name: UpdateStage3try2NextNo :exec
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
  SET status = $2
WHERE id = updated_stage3_results.try2_id
`

type UpdateStage3try2NextNoParams struct {
	ID     pgtype.UUID
	Status Stage13Status
}

// (scorer role)
func (q *Queries) UpdateStage3try2NextNo(ctx context.Context, arg UpdateStage3try2NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage3try2NextNo, arg.ID, arg.Status)
	return err
}

const updateStage3try2No1 = `-- name: UpdateStage3try2No1 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no1 = $2
WHERE id = updated_stage3_results.try2_id
RETURNING no1
`

type UpdateStage3try2No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage3try2No1(ctx context.Context, arg UpdateStage3try2No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage3try2No2 = `-- name: UpdateStage3try2No2 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no2 = $2
WHERE id = updated_stage3_results.try2_id
RETURNING no2
`

type UpdateStage3try2No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage3try2No2(ctx context.Context, arg UpdateStage3try2No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage3try2No3 = `-- name: UpdateStage3try2No3 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
) 
UPDATE stage13_tries 
SET no3 = $2 
WHERE id = updated_stage3_results.try2_id
RETURNING no3
`

type UpdateStage3try2No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage3try2No3(ctx context.Context, arg UpdateStage3try2No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}

const updateStage3try2No4 = `-- name: UpdateStage3try2No4 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no4 = $2
WHERE id = updated_stage3_results.try2_id
RETURNING no4
`

type UpdateStage3try2No4Params struct {
	ID  pgtype.UUID
	No4 string
}

// (scorer role)
func (q *Queries) UpdateStage3try2No4(ctx context.Context, arg UpdateStage3try2No4Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2No4, arg.ID, arg.No4)
	var no4 string
	err := row.Scan(&no4)
	return no4, err
}

const updateStage3try2No5 = `-- name: UpdateStage3try2No5 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no5 = $2
WHERE id = updated_stage3_results.try2_id
RETURNING no5
`

type UpdateStage3try2No5Params struct {
	ID  pgtype.UUID
	No5 string
}

// (scorer role)
func (q *Queries) UpdateStage3try2No5(ctx context.Context, arg UpdateStage3try2No5Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2No5, arg.ID, arg.No5)
	var no5 string
	err := row.Scan(&no5)
	return no5, err
}

const updateStage3try2No6 = `-- name: UpdateStage3try2No6 :one
WITH updated_stage3_results AS (
  UPDATE stage3_results
  SET
    updated_at = NOW()
  WHERE stage3_results.id = $1
  RETURNING try2_id
)
UPDATE stage13_tries
SET no6 = $2
WHERE id = updated_stage3_results.try2_id
RETURNING no6
`

type UpdateStage3try2No6Params struct {
	ID  pgtype.UUID
	No6 string
}

// (scorer role)
func (q *Queries) UpdateStage3try2No6(ctx context.Context, arg UpdateStage3try2No6Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage3try2No6, arg.ID, arg.No6)
	var no6 string
	err := row.Scan(&no6)
	return no6, err
}
