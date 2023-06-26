// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage2_try1.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage2 = `-- name: CreateStage2 :one
WITH added_stage2_try1 AS (
  INSERT INTO stage2_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), added_stage2_results AS (
  INSERT INTO stage2_results (result_id, try1_id)
  SELECT $1, id FROM added_stage2_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage2_results.id, 
  result_id, 
  status,
  no1,
  no2,
  no3,
  checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM added_stage2_try1, added_stage2_results
`

type CreateStage2Row struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Status      Stage246Status
	No1         string
	No2         string
	No3         string
	Checkmarks  string
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) CreateStage2(ctx context.Context, resultID pgtype.UUID) (CreateStage2Row, error) {
	row := q.db.QueryRow(ctx, createStage2, resultID)
	var i CreateStage2Row
	err := row.Scan(
		&i.ID,
		&i.ResultID,
		&i.Status,
		&i.No1,
		&i.No2,
		&i.No3,
		&i.Checkmarks,
		&i.IsTry2,
		&i.ShooterSign,
		&i.ScorerSign,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStage2 = `-- name: DeleteStage2 :exec
WITH deleted_stage2 AS (
  DELETE FROM stage2_results
  WHERE stage2_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage2try1 AS (
  DELETE FROM stage2_tries
  WHERE stage2_tries.id = (SELECT try1_id FROM deleted_stage2)
), deleted_stage2try2 AS (
  DELETE FROM stage2_tries
  WHERE stage2_tries.id = (SELECT try2_id FROM deleted_stage2 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '1', updated_at = NOW()
WHERE id = (SELECT result_id FROM deleted_stage2)
`

// (admin-super role)
func (q *Queries) DeleteStage2(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteStage2, id)
	return err
}

const finishStage2 = `-- name: FinishStage2 :exec
WITH get_stage2 AS (
  SELECT 
    result_id, try1_id, try2_id
  FROM stage2_results
  WHERE stage2_results.id = $1
), updated_stage2try1 AS (
  UPDATE stage2_tries
  SET status = '4'
  WHERE id = (SELECT try1_id FROM get_stage2)
), updated_stage2try2 AS (
  UPDATE stage2_tries
  SET status = '4'
  WHERE id = (SELECT try2_id FROM get_stage2 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '3', updated_at = NOW()
WHERE id = (SELECT result_id FROM get_stage2)
`

// (admin-super role)
func (q *Queries) FinishStage2(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, finishStage2, id)
	return err
}

const getStage2ById = `-- name: GetStage2ById :one
SELECT 
  stage2_results.id,
  result_id, 
  try1.status AS try1_status,
  try1.no1 AS try1_no1,
  try1.no2 AS try1_no2,
  try1.no3 AS try1_no3,
  try1.checkmarks AS try1_checkmarks,
  try2.status AS try2_status,
  try2.no1 AS try2_no1,
  try2.no2 AS try2_no2,
  try2.no3 AS try2_no3,
  try2.checkmarks AS try2_checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM stage2_results
INNER JOIN stage2_tries AS try1 ON try1.id = stage2_results.try1_id
LEFT JOIN stage2_tries AS try2 ON try2.id = stage2_results.try2_id
WHERE stage2_results.id = $1
`

type GetStage2ByIdRow struct {
	ID             pgtype.UUID
	ResultID       pgtype.UUID
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
	Try2Status     NullStage246Status
	Try2No1        sql.NullString
	Try2No2        sql.NullString
	Try2No3        sql.NullString
	Try2Checkmarks sql.NullString
	IsTry2         bool
	ShooterSign    pgtype.Text
	ScorerSign     pgtype.Text
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

// (all role)
func (q *Queries) GetStage2ById(ctx context.Context, id pgtype.UUID) (GetStage2ByIdRow, error) {
	row := q.db.QueryRow(ctx, getStage2ById, id)
	var i GetStage2ByIdRow
	err := row.Scan(
		&i.ID,
		&i.ResultID,
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
		&i.IsTry2,
		&i.ShooterSign,
		&i.ScorerSign,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStage2RelationByResultId = `-- name: GetStage2RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage2_results
WHERE result_id = $1
`

type GetStage2RelationByResultIdRow struct {
	ID       pgtype.UUID
	ResultID pgtype.UUID
	IsTry2   bool
}

// (all role)
func (q *Queries) GetStage2RelationByResultId(ctx context.Context, resultID pgtype.UUID) (GetStage2RelationByResultIdRow, error) {
	row := q.db.QueryRow(ctx, getStage2RelationByResultId, resultID)
	var i GetStage2RelationByResultIdRow
	err := row.Scan(&i.ID, &i.ResultID, &i.IsTry2)
	return i, err
}

const getStage2try1Status = `-- name: GetStage2try1Status :one
SELECT 
  status
FROM stage2_results
INNER JOIN stage2_tries ON stage2_tries.id = stage2_results.try1_id
WHERE stage2_results.id = $1
`

func (q *Queries) GetStage2try1Status(ctx context.Context, id pgtype.UUID) (Stage246Status, error) {
	row := q.db.QueryRow(ctx, getStage2try1Status, id)
	var status Stage246Status
	err := row.Scan(&status)
	return status, err
}

const getStage2try2ExistById = `-- name: GetStage2try2ExistById :one
SELECT 
  is_try2
FROM stage2_results
WHERE id = $1
`

// (all role)
func (q *Queries) GetStage2try2ExistById(ctx context.Context, id pgtype.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, getStage2try2ExistById, id)
	var is_try2 bool
	err := row.Scan(&is_try2)
	return is_try2, err
}

const updateStage2NextTry = `-- name: UpdateStage2NextTry :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET status = '4'
WHERE id = (SELECT try1_id FROM stage2_results)
`

// (scorer role)
func (q *Queries) UpdateStage2NextTry(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, updateStage2NextTry, id)
	return err
}

const updateStage2Signs = `-- name: UpdateStage2Signs :one
UPDATE stage2_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at
`

type UpdateStage2SignsParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

type UpdateStage2SignsRow struct {
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	UpdatedAt   pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage2Signs(ctx context.Context, arg UpdateStage2SignsParams) (UpdateStage2SignsRow, error) {
	row := q.db.QueryRow(ctx, updateStage2Signs, arg.ID, arg.ShooterSign, arg.ScorerSign)
	var i UpdateStage2SignsRow
	err := row.Scan(&i.ShooterSign, &i.ScorerSign, &i.UpdatedAt)
	return i, err
}

const updateStage2try1 = `-- name: UpdateStage2try1 :one
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
)
SELECT 
  updated_stage2_try1.status AS try1_status,
  updated_stage2_try1.no1 AS try1_no1,
  updated_stage2_try1.no2 AS try1_no2,
  updated_stage2_try1.no3 AS try1_no3,
  updated_stage2_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage2_results, updated_stage2_try1
`

type UpdateStage2try1Params struct {
	ID             pgtype.UUID
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
}

type UpdateStage2try1Row struct {
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
	UpdatedAt      pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage2try1(ctx context.Context, arg UpdateStage2try1Params) (UpdateStage2try1Row, error) {
	row := q.db.QueryRow(ctx, updateStage2try1,
		arg.ID,
		arg.Try1Status,
		arg.Try1No1,
		arg.Try1No2,
		arg.Try1No3,
		arg.Try1Checkmarks,
	)
	var i UpdateStage2try1Row
	err := row.Scan(
		&i.Try1Status,
		&i.Try1No1,
		&i.Try1No2,
		&i.Try1No3,
		&i.Try1Checkmarks,
		&i.UpdatedAt,
	)
	return i, err
}

const updateStage2try1Checkmarks = `-- name: UpdateStage2try1Checkmarks :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try1_id FROM stage2_results)
RETURNING checkmarks
`

type UpdateStage2try1CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage2try1Checkmarks(ctx context.Context, arg UpdateStage2try1CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try1Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage2try1FinishFailed = `-- name: UpdateStage2try1FinishFailed :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try1_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '3'
  WHERE id = (SELECT try1_id FROM updated_stage2_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage2_results)
`

type UpdateStage2try1FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage2try1FinishFailed(ctx context.Context, arg UpdateStage2try1FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage2try1FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage2try1FinishSuccess = `-- name: UpdateStage2try1FinishSuccess :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage2_results.id = $1
  RETURNING result_id, try1_id
), updated_stage2_tries AS (
  UPDATE stage2_tries
    SET status = '4'
  WHERE id = (SELECT try1_id FROM updated_stage2_results)
)
UPDATE results 
SET stage = '3', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage2_results)
`

type UpdateStage2try1FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage2try1FinishSuccess(ctx context.Context, arg UpdateStage2try1FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage2try1FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage2try1NextNo = `-- name: UpdateStage2try1NextNo :exec
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
  SET status = $2
WHERE id = (SELECT try1_id FROM stage2_results)
`

type UpdateStage2try1NextNoParams struct {
	ID     pgtype.UUID
	Status Stage246Status
}

// (scorer role)
func (q *Queries) UpdateStage2try1NextNo(ctx context.Context, arg UpdateStage2try1NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage2try1NextNo, arg.ID, arg.Status)
	return err
}

const updateStage2try1No1 = `-- name: UpdateStage2try1No1 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET no1 = $2
WHERE id = (SELECT try1_id FROM stage2_results)
RETURNING no1
`

type UpdateStage2try1No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage2try1No1(ctx context.Context, arg UpdateStage2try1No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try1No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage2try1No2 = `-- name: UpdateStage2try1No2 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries 
SET no2 = $2
WHERE id = (SELECT try1_id FROM stage2_results)
RETURNING no2
`

type UpdateStage2try1No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage2try1No2(ctx context.Context, arg UpdateStage2try1No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try1No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage2try1No3 = `-- name: UpdateStage2try1No3 :one
WITH updated_stage2_results AS (
  UPDATE stage2_results
  SET
    updated_at = NOW()
  WHERE stage2_results.id = $1
  RETURNING try1_id
)
UPDATE stage2_tries
SET no3 = $2
WHERE id = (SELECT try1_id FROM stage2_results) 
RETURNING no3
`

type UpdateStage2try1No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage2try1No3(ctx context.Context, arg UpdateStage2try1No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage2try1No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}
