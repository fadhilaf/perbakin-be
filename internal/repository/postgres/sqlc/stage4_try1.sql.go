// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage4_try1.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage4 = `-- name: CreateStage4 :one
WITH added_stage4_try1 AS (
  INSERT INTO stage46_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), added_stage4_results AS (
  INSERT INTO stage4_results (result_id, try1_id)
  SELECT $1, id FROM added_stage4_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage4_results.id, 
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
FROM added_stage4_try1, added_stage4_results
`

type CreateStage4Row struct {
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

func (q *Queries) CreateStage4(ctx context.Context, resultID pgtype.UUID) (CreateStage4Row, error) {
	row := q.db.QueryRow(ctx, createStage4, resultID)
	var i CreateStage4Row
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

const deleteStage4 = `-- name: DeleteStage4 :exec
WITH deleted_stage4 AS (
  DELETE FROM stage4_results
  WHERE stage4_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage4try1 AS (
  DELETE FROM stage46_tries
  WHERE stage46_tries.id IN (SELECT try1_id FROM deleted_stage4)
)
DELETE FROM stage46_tries
WHERE stage46_tries.id = (SELECT try2_id FROM deleted_stage4 WHERE try2_id IS NOT NULL)
`

// (admin-super role)
func (q *Queries) DeleteStage4(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteStage4, id)
	return err
}

const finishStage4 = `-- name: FinishStage4 :exec
WITH get_stage4 AS (
  SELECT 
    result_id, try1_id, try2_id
  FROM stage4_results
  WHERE stage4_results.id = $1
), updated_stage4try1 AS (
  UPDATE stage46_tries
  SET status = '4'
  WHERE id = (SELECT try1_id FROM get_stage4)
), updated_stage4try2 AS (
  UPDATE stage46_tries
  SET status = '4'
  WHERE id = (SELECT try2_id FROM get_stage4 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '5', updated_at = NOW()
WHERE id = (SELECT result_id FROM get_stage4)
`

// (admin-super role)
func (q *Queries) FinishStage4(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, finishStage4, id)
	return err
}

const getStage4ById = `-- name: GetStage4ById :one
SELECT 
  stage4_results.id,
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
FROM stage4_results
INNER JOIN stage46_tries AS try1 ON try1.id = stage4_results.try1_id
LEFT JOIN stage46_tries AS try2 ON try2.id = stage4_results.try2_id
WHERE stage4_results.id = $1
`

type GetStage4ByIdRow struct {
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
func (q *Queries) GetStage4ById(ctx context.Context, id pgtype.UUID) (GetStage4ByIdRow, error) {
	row := q.db.QueryRow(ctx, getStage4ById, id)
	var i GetStage4ByIdRow
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

const getStage4RelationByResultId = `-- name: GetStage4RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage4_results
WHERE result_id = $1
`

type GetStage4RelationByResultIdRow struct {
	ID       pgtype.UUID
	ResultID pgtype.UUID
	IsTry2   bool
}

// (all role)
func (q *Queries) GetStage4RelationByResultId(ctx context.Context, resultID pgtype.UUID) (GetStage4RelationByResultIdRow, error) {
	row := q.db.QueryRow(ctx, getStage4RelationByResultId, resultID)
	var i GetStage4RelationByResultIdRow
	err := row.Scan(&i.ID, &i.ResultID, &i.IsTry2)
	return i, err
}

const getStage4try1Status = `-- name: GetStage4try1Status :one
SELECT 
  status
FROM stage4_results
INNER JOIN stage46_tries ON stage46_tries.id = stage4_results.try1_id
WHERE stage4_results.id = $1
`

func (q *Queries) GetStage4try1Status(ctx context.Context, id pgtype.UUID) (Stage246Status, error) {
	row := q.db.QueryRow(ctx, getStage4try1Status, id)
	var status Stage246Status
	err := row.Scan(&status)
	return status, err
}

const getStage4try2ExistById = `-- name: GetStage4try2ExistById :one
SELECT 
  is_try2
FROM stage4_results
WHERE id = $1
`

// (all role)
func (q *Queries) GetStage4try2ExistById(ctx context.Context, id pgtype.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, getStage4try2ExistById, id)
	var is_try2 bool
	err := row.Scan(&is_try2)
	return is_try2, err
}

const updateStage4NextTry = `-- name: UpdateStage4NextTry :exec
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET status = '4'
FROM updated_stage4_results
WHERE id = updated_stage4_results.try1_id
`

// (scorer role)
func (q *Queries) UpdateStage4NextTry(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, updateStage4NextTry, id)
	return err
}

const updateStage4Signs = `-- name: UpdateStage4Signs :one
UPDATE stage4_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at
`

type UpdateStage4SignsParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

type UpdateStage4SignsRow struct {
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	UpdatedAt   pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage4Signs(ctx context.Context, arg UpdateStage4SignsParams) (UpdateStage4SignsRow, error) {
	row := q.db.QueryRow(ctx, updateStage4Signs, arg.ID, arg.ShooterSign, arg.ScorerSign)
	var i UpdateStage4SignsRow
	err := row.Scan(&i.ShooterSign, &i.ScorerSign, &i.UpdatedAt)
	return i, err
}

const updateStage4try1 = `-- name: UpdateStage4try1 :one
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage4_try1 AS (
  UPDATE stage46_tries
  SET 
    status = $2,
    no1 = $3,
    no2 = $4,
    no3 = $5,
    checkmarks = $6
  FROM updated_stage4_results
  WHERE id = updated_stage4_results.try1_id
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
)
SELECT 
  updated_stage4_try1.status AS try1_status,
  updated_stage4_try1.no1 AS try1_no1,
  updated_stage4_try1.no2 AS try1_no2,
  updated_stage4_try1.no3 AS try1_no3,
  updated_stage4_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage4_results, updated_stage4_try1
`

type UpdateStage4try1Params struct {
	ID             pgtype.UUID
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
}

type UpdateStage4try1Row struct {
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
	UpdatedAt      pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage4try1(ctx context.Context, arg UpdateStage4try1Params) (UpdateStage4try1Row, error) {
	row := q.db.QueryRow(ctx, updateStage4try1,
		arg.ID,
		arg.Try1Status,
		arg.Try1No1,
		arg.Try1No2,
		arg.Try1No3,
		arg.Try1Checkmarks,
	)
	var i UpdateStage4try1Row
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

const updateStage4try1Checkmarks = `-- name: UpdateStage4try1Checkmarks :one
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET 
  checkmarks = $2
FROM updated_stage4_results
WHERE id = updated_stage4_results.try1_id
RETURNING checkmarks
`

type UpdateStage4try1CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage4try1Checkmarks(ctx context.Context, arg UpdateStage4try1CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage4try1Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage4try1FinishFailed = `-- name: UpdateStage4try1FinishFailed :exec
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage4_results.id = $1
  RETURNING result_id, try1_id
), updated_stage4_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  FROM updated_stage4_results
  WHERE id = updated_stage4_results.try1_id
)
UPDATE results 
SET failed = true, updated_at = NOW()
FROM updated_stage4_results
WHERE id = updated_stage4_results.result_id
`

type UpdateStage4try1FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage4try1FinishFailed(ctx context.Context, arg UpdateStage4try1FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage4try1FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage4try1FinishSuccess = `-- name: UpdateStage4try1FinishSuccess :exec
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage4_results.id = $1
  RETURNING result_id, try1_id
), updated_stage4_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  FROM updated_stage4_results
  WHERE id = updated_stage4_results.try1_id
)
UPDATE results 
SET stage = '5', updated_at = NOW()
FROM updated_stage4_results
WHERE id = updated_stage4_results.result_id
`

type UpdateStage4try1FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage4try1FinishSuccess(ctx context.Context, arg UpdateStage4try1FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage4try1FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage4try1NextNo = `-- name: UpdateStage4try1NextNo :exec
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
  SET status = $2
FROM updated_stage4_results
WHERE id = updated_stage4_results.try1_id
`

type UpdateStage4try1NextNoParams struct {
	ID     pgtype.UUID
	Status Stage246Status
}

// (scorer role)
func (q *Queries) UpdateStage4try1NextNo(ctx context.Context, arg UpdateStage4try1NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage4try1NextNo, arg.ID, arg.Status)
	return err
}

const updateStage4try1No1 = `-- name: UpdateStage4try1No1 :one
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET no1 = $2
FROM updated_stage4_results
WHERE id = updated_stage4_results.try1_id
RETURNING no1
`

type UpdateStage4try1No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage4try1No1(ctx context.Context, arg UpdateStage4try1No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage4try1No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage4try1No2 = `-- name: UpdateStage4try1No2 :one
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries 
SET no2 = $2
FROM updated_stage4_results
WHERE id = updated_stage4_results.try1_id
RETURNING no2
`

type UpdateStage4try1No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage4try1No2(ctx context.Context, arg UpdateStage4try1No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage4try1No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage4try1No3 = `-- name: UpdateStage4try1No3 :one
WITH updated_stage4_results AS (
  UPDATE stage4_results
  SET
    updated_at = NOW()
  WHERE stage4_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET no3 = $2
FROM updated_stage4_results
WHERE id = updated_stage4_results.try1_id 
RETURNING no3
`

type UpdateStage4try1No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage4try1No3(ctx context.Context, arg UpdateStage4try1No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage4try1No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}
