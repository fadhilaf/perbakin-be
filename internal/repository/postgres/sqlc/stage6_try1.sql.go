// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage6_try1.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage6 = `-- name: CreateStage6 :one
WITH added_stage6_try1 AS (
  INSERT INTO stage46_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, checkmarks
), added_stage6_results AS (
  INSERT INTO stage6_results (result_id, try1_id)
  SELECT $1, id FROM added_stage6_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage6_results.id, 
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
FROM added_stage6_try1, added_stage6_results
`

type CreateStage6Row struct {
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

func (q *Queries) CreateStage6(ctx context.Context, resultID pgtype.UUID) (CreateStage6Row, error) {
	row := q.db.QueryRow(ctx, createStage6, resultID)
	var i CreateStage6Row
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

const deleteStage6 = `-- name: DeleteStage6 :exec
WITH deleted_stage6 AS (
  DELETE FROM stage6_results
  WHERE stage6_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage6try1 AS (
  DELETE FROM stage46_tries
  WHERE stage46_tries.id IN (SELECT try1_id FROM deleted_stage6)
)
DELETE FROM stage46_tries
WHERE stage46_tries.id = (SELECT try2_id FROM deleted_stage6 WHERE try2_id IS NOT NULL)
`

// (admin-super role)
func (q *Queries) DeleteStage6(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteStage6, id)
	return err
}

const finishStage6 = `-- name: FinishStage6 :exec
WITH get_stage6 AS (
  SELECT 
    result_id, try1_id, try2_id
  FROM stage6_results
  WHERE stage6_results.id = $1
), updated_stage6try1 AS (
  UPDATE stage46_tries
  SET status = '4'
  FROM get_stage6
  WHERE id = get_stage6.try1_id
), updated_stage6try2 AS (
  UPDATE stage46_tries
  SET status = '4'
  WHERE id IN (SELECT try2_id FROM get_stage6 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '7', updated_at = NOW()
FROM get_stage6
WHERE id = get_stage6.result_id
`

// (admin-super role)
func (q *Queries) FinishStage6(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, finishStage6, id)
	return err
}

const getStage6ById = `-- name: GetStage6ById :one
SELECT 
  stage6_results.id,
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
FROM stage6_results
INNER JOIN stage46_tries AS try1 ON try1.id = stage6_results.try1_id
LEFT JOIN stage46_tries AS try2 ON try2.id = stage6_results.try2_id
WHERE stage6_results.id = $1
`

type GetStage6ByIdRow struct {
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
func (q *Queries) GetStage6ById(ctx context.Context, id pgtype.UUID) (GetStage6ByIdRow, error) {
	row := q.db.QueryRow(ctx, getStage6ById, id)
	var i GetStage6ByIdRow
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

const getStage6RelationByResultId = `-- name: GetStage6RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage6_results
WHERE result_id = $1
`

type GetStage6RelationByResultIdRow struct {
	ID       pgtype.UUID
	ResultID pgtype.UUID
	IsTry2   bool
}

// (all role)
func (q *Queries) GetStage6RelationByResultId(ctx context.Context, resultID pgtype.UUID) (GetStage6RelationByResultIdRow, error) {
	row := q.db.QueryRow(ctx, getStage6RelationByResultId, resultID)
	var i GetStage6RelationByResultIdRow
	err := row.Scan(&i.ID, &i.ResultID, &i.IsTry2)
	return i, err
}

const getStage6try1Status = `-- name: GetStage6try1Status :one
SELECT 
  status
FROM stage6_results
INNER JOIN stage46_tries ON stage46_tries.id = stage6_results.try1_id
WHERE stage6_results.id = $1
`

func (q *Queries) GetStage6try1Status(ctx context.Context, id pgtype.UUID) (Stage246Status, error) {
	row := q.db.QueryRow(ctx, getStage6try1Status, id)
	var status Stage246Status
	err := row.Scan(&status)
	return status, err
}

const getStage6try2ExistById = `-- name: GetStage6try2ExistById :one
SELECT 
  is_try2
FROM stage6_results
WHERE id = $1
`

// (all role)
func (q *Queries) GetStage6try2ExistById(ctx context.Context, id pgtype.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, getStage6try2ExistById, id)
	var is_try2 bool
	err := row.Scan(&is_try2)
	return is_try2, err
}

const updateStage6NextTry = `-- name: UpdateStage6NextTry :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET status = '4'
FROM updated_stage6_results
WHERE id = updated_stage6_results.try1_id
`

// (scorer role)
func (q *Queries) UpdateStage6NextTry(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, updateStage6NextTry, id)
	return err
}

const updateStage6Signs = `-- name: UpdateStage6Signs :one
UPDATE stage6_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at
`

type UpdateStage6SignsParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

type UpdateStage6SignsRow struct {
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	UpdatedAt   pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage6Signs(ctx context.Context, arg UpdateStage6SignsParams) (UpdateStage6SignsRow, error) {
	row := q.db.QueryRow(ctx, updateStage6Signs, arg.ID, arg.ShooterSign, arg.ScorerSign)
	var i UpdateStage6SignsRow
	err := row.Scan(&i.ShooterSign, &i.ScorerSign, &i.UpdatedAt)
	return i, err
}

const updateStage6try1 = `-- name: UpdateStage6try1 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage6_try1 AS (
  UPDATE stage46_tries
  SET 
    status = $2,
    no1 = $3,
    no2 = $4,
    no3 = $5,
    checkmarks = $6
  FROM updated_stage6_results
  WHERE id = updated_stage6_results.try1_id
  RETURNING 
    status,
    no1,
    no2,
    no3,
    checkmarks
)
SELECT 
  updated_stage6_try1.status AS try1_status,
  updated_stage6_try1.no1 AS try1_no1,
  updated_stage6_try1.no2 AS try1_no2,
  updated_stage6_try1.no3 AS try1_no3,
  updated_stage6_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage6_results, updated_stage6_try1
`

type UpdateStage6try1Params struct {
	ID             pgtype.UUID
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
}

type UpdateStage6try1Row struct {
	Try1Status     Stage246Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1Checkmarks string
	UpdatedAt      pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage6try1(ctx context.Context, arg UpdateStage6try1Params) (UpdateStage6try1Row, error) {
	row := q.db.QueryRow(ctx, updateStage6try1,
		arg.ID,
		arg.Try1Status,
		arg.Try1No1,
		arg.Try1No2,
		arg.Try1No3,
		arg.Try1Checkmarks,
	)
	var i UpdateStage6try1Row
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

const updateStage6try1Checkmarks = `-- name: UpdateStage6try1Checkmarks :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET 
  checkmarks = $2
FROM updated_stage6_results
WHERE id = updated_stage6_results.try1_id
RETURNING checkmarks
`

type UpdateStage6try1CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage6try1Checkmarks(ctx context.Context, arg UpdateStage6try1CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage6try1Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage6try1FinishFailed = `-- name: UpdateStage6try1FinishFailed :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage6_results.id = $1
  RETURNING result_id, try1_id
), updated_stage6_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  FROM updated_stage6_results
  WHERE id = updated_stage6_results.try1_id
)
UPDATE results 
SET failed = true, updated_at = NOW()
FROM updated_stage6_results
WHERE id = updated_stage6_results.result_id
`

type UpdateStage6try1FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage6try1FinishFailed(ctx context.Context, arg UpdateStage6try1FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage6try1FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage6try1FinishSuccess = `-- name: UpdateStage6try1FinishSuccess :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage6_results.id = $1
  RETURNING result_id, try1_id
), updated_stage6_tries AS (
  UPDATE stage46_tries
    SET status = '4'
  FROM updated_stage6_results
  WHERE id = updated_stage6_results.try1_id
)
UPDATE results 
SET stage = '7', updated_at = NOW()
FROM updated_stage6_results
WHERE id = updated_stage6_results.result_id
`

type UpdateStage6try1FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage6try1FinishSuccess(ctx context.Context, arg UpdateStage6try1FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage6try1FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage6try1NextNo = `-- name: UpdateStage6try1NextNo :exec
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
  SET status = $2
FROM updated_stage6_results
WHERE id = updated_stage6_results.try1_id
`

type UpdateStage6try1NextNoParams struct {
	ID     pgtype.UUID
	Status Stage246Status
}

// (scorer role)
func (q *Queries) UpdateStage6try1NextNo(ctx context.Context, arg UpdateStage6try1NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage6try1NextNo, arg.ID, arg.Status)
	return err
}

const updateStage6try1No1 = `-- name: UpdateStage6try1No1 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET no1 = $2
FROM updated_stage6_results
WHERE id = updated_stage6_results.try1_id
RETURNING no1
`

type UpdateStage6try1No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage6try1No1(ctx context.Context, arg UpdateStage6try1No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage6try1No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage6try1No2 = `-- name: UpdateStage6try1No2 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries 
SET no2 = $2
FROM updated_stage6_results
WHERE id = updated_stage6_results.try1_id
RETURNING no2
`

type UpdateStage6try1No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage6try1No2(ctx context.Context, arg UpdateStage6try1No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage6try1No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage6try1No3 = `-- name: UpdateStage6try1No3 :one
WITH updated_stage6_results AS (
  UPDATE stage6_results
  SET
    updated_at = NOW()
  WHERE stage6_results.id = $1
  RETURNING try1_id
)
UPDATE stage46_tries
SET no3 = $2
FROM updated_stage6_results
WHERE id = updated_stage6_results.try1_id 
RETURNING no3
`

type UpdateStage6try1No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage6try1No3(ctx context.Context, arg UpdateStage6try1No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage6try1No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}
