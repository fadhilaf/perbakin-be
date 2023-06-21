// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: stage1_try1.sql

package postgres

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStage1 = `-- name: CreateStage1 :one
WITH added_stage1_try1 AS (
  INSERT INTO stage13_tries DEFAULT VALUES
  RETURNING id, status, no1, no2, no3, no4, no5, no6, checkmarks
), added_stage1_results AS (
  INSERT INTO stage1_results (result_id, try1_id)
  SELECT $1, id FROM added_stage1_try1
  RETURNING id, result_id, try1_id, is_try2, shooter_sign, scorer_sign, created_at, updated_at
)
SELECT
  added_stage1_results.id, 
  result_id, 
  status,
  no1,
  no2,
  no3,
  no4,
  no5,
  no6,
  checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM added_stage1_try1, added_stage1_results
`

type CreateStage1Row struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Status      Stage13Status
	No1         string
	No2         string
	No3         string
	No4         string
	No5         string
	No6         string
	Checkmarks  string
	IsTry2      bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) CreateStage1(ctx context.Context, resultID pgtype.UUID) (CreateStage1Row, error) {
	row := q.db.QueryRow(ctx, createStage1, resultID)
	var i CreateStage1Row
	err := row.Scan(
		&i.ID,
		&i.ResultID,
		&i.Status,
		&i.No1,
		&i.No2,
		&i.No3,
		&i.No4,
		&i.No5,
		&i.No6,
		&i.Checkmarks,
		&i.IsTry2,
		&i.ShooterSign,
		&i.ScorerSign,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStage1 = `-- name: DeleteStage1 :exec
WITH deleted_stage1 AS (
  DELETE FROM stage1_results
  WHERE stage1_results.id = $1
  RETURNING result_id, try1_id, try2_id
), deleted_stage1try1 AS (
  DELETE FROM stage13_tries
  WHERE stage13_tries.id = (SELECT try1_id FROM deleted_stage1)
), deleted_stage1try2 AS (
  DELETE FROM stage13_tries
  WHERE stage13_tries.id = (SELECT try2_id FROM deleted_stage1 WHERE try2_id IS NOT NULL)
)
UPDATE results 
SET stage = '1', updated_at = NOW()
WHERE id = (SELECT result_id FROM deleted_stage1)
`

// (admin-super role)
func (q *Queries) DeleteStage1(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteStage1, id)
	return err
}

const getStage1ById = `-- name: GetStage1ById :one
SELECT 
  stage1_results.id,
  result_id, 
  try1.status AS try1_status,
  try1.no1 AS try1_no1,
  try1.no2 AS try1_no2,
  try1.no3 AS try1_no3,
  try1.no4 AS try1_no4,
  try1.no5 AS try1_no5,
  try1.no6 AS try1_no6,
  try1.checkmarks AS try1_checkmarks,
  try2.status AS try2_status,
  try2.no1 AS try2_no1,
  try2.no2 AS try2_no2,
  try2.no3 AS try2_no3,
  try2.no4 AS try2_no4,
  try2.no5 AS try2_no5,
  try2.no6 AS try2_no6,
  try2.checkmarks AS try2_checkmarks,
  is_try2,
  shooter_sign,
  scorer_sign,
  created_at,
  updated_at
FROM stage1_results
INNER JOIN stage13_tries AS try1 ON try1.id = stage1_results.try1_id
LEFT JOIN stage13_tries AS try2 ON try2.id = stage1_results.try2_id
WHERE stage1_results.id = $1
`

type GetStage1ByIdRow struct {
	ID             pgtype.UUID
	ResultID       pgtype.UUID
	Try1Status     Stage13Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1No4        string
	Try1No5        string
	Try1No6        string
	Try1Checkmarks string
	Try2Status     NullStage13Status
	Try2No1        sql.NullString
	Try2No2        sql.NullString
	Try2No3        sql.NullString
	Try2No4        sql.NullString
	Try2No5        sql.NullString
	Try2No6        sql.NullString
	Try2Checkmarks sql.NullString
	IsTry2         bool
	ShooterSign    pgtype.Text
	ScorerSign     pgtype.Text
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

// (all role)
func (q *Queries) GetStage1ById(ctx context.Context, id pgtype.UUID) (GetStage1ByIdRow, error) {
	row := q.db.QueryRow(ctx, getStage1ById, id)
	var i GetStage1ByIdRow
	err := row.Scan(
		&i.ID,
		&i.ResultID,
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
		&i.IsTry2,
		&i.ShooterSign,
		&i.ScorerSign,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getStage1RelationByResultId = `-- name: GetStage1RelationByResultId :one
SELECT 
  id, 
  result_id,
  is_try2
FROM stage1_results
WHERE result_id = $1
`

type GetStage1RelationByResultIdRow struct {
	ID       pgtype.UUID
	ResultID pgtype.UUID
	IsTry2   bool
}

// (all role)
func (q *Queries) GetStage1RelationByResultId(ctx context.Context, resultID pgtype.UUID) (GetStage1RelationByResultIdRow, error) {
	row := q.db.QueryRow(ctx, getStage1RelationByResultId, resultID)
	var i GetStage1RelationByResultIdRow
	err := row.Scan(&i.ID, &i.ResultID, &i.IsTry2)
	return i, err
}

const getStage1try1Status = `-- name: GetStage1try1Status :one
SELECT 
  status
FROM stage1_results
INNER JOIN stage13_tries ON stage13_tries.id = stage1_results.try1_id
WHERE stage1_results.id = $1
`

func (q *Queries) GetStage1try1Status(ctx context.Context, id pgtype.UUID) (Stage13Status, error) {
	row := q.db.QueryRow(ctx, getStage1try1Status, id)
	var status Stage13Status
	err := row.Scan(&status)
	return status, err
}

const getStage1try2ExistById = `-- name: GetStage1try2ExistById :one
SELECT 
  is_try2
FROM stage1_results
WHERE id = $1
`

// (all role)
func (q *Queries) GetStage1try2ExistById(ctx context.Context, id pgtype.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, getStage1try2ExistById, id)
	var is_try2 bool
	err := row.Scan(&is_try2)
	return is_try2, err
}

const updateStage1NextTry = `-- name: UpdateStage1NextTry :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET 
    is_try2 = true,
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET status = '7'
WHERE id = (SELECT try1_id FROM stage1_results)
`

// (scorer role)
func (q *Queries) UpdateStage1NextTry(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, updateStage1NextTry, id)
	return err
}

const updateStage1Signs = `-- name: UpdateStage1Signs :one
UPDATE stage1_results 
SET 
  shooter_sign = $2, 
  scorer_sign = $3, 
  updated_at = NOW()
WHERE id = $1
RETURNING shooter_sign, scorer_sign, updated_at
`

type UpdateStage1SignsParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

type UpdateStage1SignsRow struct {
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
	UpdatedAt   pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage1Signs(ctx context.Context, arg UpdateStage1SignsParams) (UpdateStage1SignsRow, error) {
	row := q.db.QueryRow(ctx, updateStage1Signs, arg.ID, arg.ShooterSign, arg.ScorerSign)
	var i UpdateStage1SignsRow
	err := row.Scan(&i.ShooterSign, &i.ScorerSign, &i.UpdatedAt)
	return i, err
}

const updateStage1try1 = `-- name: UpdateStage1try1 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1 
  RETURNING try1_id, try2_id, is_try2, updated_at
), updated_stage1_try1 AS (
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
  WHERE id = (SELECT try1_id FROM stage1_results)
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
  updated_stage1_try1.status AS try1_status,
  updated_stage1_try1.no1 AS try1_no1,
  updated_stage1_try1.no2 AS try1_no2,
  updated_stage1_try1.no3 AS try1_no3,
  updated_stage1_try1.no4 AS try1_no4,
  updated_stage1_try1.no5 AS try1_no5,
  updated_stage1_try1.no6 AS try1_no6,
  updated_stage1_try1.checkmarks AS try1_checkmarks,
  updated_at
FROM updated_stage1_results, updated_stage1_try1
`

type UpdateStage1try1Params struct {
	ID             pgtype.UUID
	Try1Status     Stage13Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1No4        string
	Try1No5        string
	Try1No6        string
	Try1Checkmarks string
}

type UpdateStage1try1Row struct {
	Try1Status     Stage13Status
	Try1No1        string
	Try1No2        string
	Try1No3        string
	Try1No4        string
	Try1No5        string
	Try1No6        string
	Try1Checkmarks string
	UpdatedAt      pgtype.Timestamp
}

// (admin-super role)
func (q *Queries) UpdateStage1try1(ctx context.Context, arg UpdateStage1try1Params) (UpdateStage1try1Row, error) {
	row := q.db.QueryRow(ctx, updateStage1try1,
		arg.ID,
		arg.Try1Status,
		arg.Try1No1,
		arg.Try1No2,
		arg.Try1No3,
		arg.Try1No4,
		arg.Try1No5,
		arg.Try1No6,
		arg.Try1Checkmarks,
	)
	var i UpdateStage1try1Row
	err := row.Scan(
		&i.Try1Status,
		&i.Try1No1,
		&i.Try1No2,
		&i.Try1No3,
		&i.Try1No4,
		&i.Try1No5,
		&i.Try1No6,
		&i.Try1Checkmarks,
		&i.UpdatedAt,
	)
	return i, err
}

const updateStage1try1Checkmarks = `-- name: UpdateStage1try1Checkmarks :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET 
  checkmarks = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING checkmarks
`

type UpdateStage1try1CheckmarksParams struct {
	ID         pgtype.UUID
	Checkmarks string
}

// (scorer role)
func (q *Queries) UpdateStage1try1Checkmarks(ctx context.Context, arg UpdateStage1try1CheckmarksParams) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1Checkmarks, arg.ID, arg.Checkmarks)
	var checkmarks string
	err := row.Scan(&checkmarks)
	return checkmarks, err
}

const updateStage1try1FinishFailed = `-- name: UpdateStage1try1FinishFailed :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try1_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try1_id FROM updated_stage1_results)
)
UPDATE results 
SET failed = true, updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results)
`

type UpdateStage1try1FinishFailedParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage1try1FinishFailed(ctx context.Context, arg UpdateStage1try1FinishFailedParams) error {
	_, err := q.db.Exec(ctx, updateStage1try1FinishFailed, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage1try1FinishSuccess = `-- name: UpdateStage1try1FinishSuccess :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW(), 
  	shooter_sign = $2, 
  	scorer_sign = $3
  WHERE stage1_results.id = $1
  RETURNING result_id, try1_id
), updated_stage13_tries AS (
  UPDATE stage13_tries
    SET status = '7'
  WHERE id = (SELECT try1_id FROM updated_stage1_results)
)
UPDATE results 
SET stage = '2', updated_at = NOW()
WHERE id = (SELECT result_id FROM updated_stage1_results)
`

type UpdateStage1try1FinishSuccessParams struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

// (scorer role)
func (q *Queries) UpdateStage1try1FinishSuccess(ctx context.Context, arg UpdateStage1try1FinishSuccessParams) error {
	_, err := q.db.Exec(ctx, updateStage1try1FinishSuccess, arg.ID, arg.ShooterSign, arg.ScorerSign)
	return err
}

const updateStage1try1NextNo = `-- name: UpdateStage1try1NextNo :exec
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
  SET status = $2
WHERE id = (SELECT try1_id FROM stage1_results)
`

type UpdateStage1try1NextNoParams struct {
	ID     pgtype.UUID
	Status Stage13Status
}

// (scorer role)
func (q *Queries) UpdateStage1try1NextNo(ctx context.Context, arg UpdateStage1try1NextNoParams) error {
	_, err := q.db.Exec(ctx, updateStage1try1NextNo, arg.ID, arg.Status)
	return err
}

const updateStage1try1No1 = `-- name: UpdateStage1try1No1 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET no1 = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING no1
`

type UpdateStage1try1No1Params struct {
	ID  pgtype.UUID
	No1 string
}

// (scorer role)
func (q *Queries) UpdateStage1try1No1(ctx context.Context, arg UpdateStage1try1No1Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1No1, arg.ID, arg.No1)
	var no1 string
	err := row.Scan(&no1)
	return no1, err
}

const updateStage1try1No2 = `-- name: UpdateStage1try1No2 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries 
SET no2 = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING no2
`

type UpdateStage1try1No2Params struct {
	ID  pgtype.UUID
	No2 string
}

// (scorer role)
func (q *Queries) UpdateStage1try1No2(ctx context.Context, arg UpdateStage1try1No2Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1No2, arg.ID, arg.No2)
	var no2 string
	err := row.Scan(&no2)
	return no2, err
}

const updateStage1try1No3 = `-- name: UpdateStage1try1No3 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET no3 = $2
WHERE id = (SELECT try1_id FROM stage1_results) 
RETURNING no3
`

type UpdateStage1try1No3Params struct {
	ID  pgtype.UUID
	No3 string
}

// (scorer role)
func (q *Queries) UpdateStage1try1No3(ctx context.Context, arg UpdateStage1try1No3Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1No3, arg.ID, arg.No3)
	var no3 string
	err := row.Scan(&no3)
	return no3, err
}

const updateStage1try1No4 = `-- name: UpdateStage1try1No4 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
)
UPDATE stage13_tries
SET no4 = $2
WHERE id = (SELECT try1_id FROM stage1_results) 
RETURNING no4
`

type UpdateStage1try1No4Params struct {
	ID  pgtype.UUID
	No4 string
}

// (scorer role)
func (q *Queries) UpdateStage1try1No4(ctx context.Context, arg UpdateStage1try1No4Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1No4, arg.ID, arg.No4)
	var no4 string
	err := row.Scan(&no4)
	return no4, err
}

const updateStage1try1No5 = `-- name: UpdateStage1try1No5 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1
  RETURNING try1_id
) 
UPDATE stage13_tries 
SET no5 = $2
WHERE id = (SELECT try1_id FROM stage1_results) 
RETURNING no5
`

type UpdateStage1try1No5Params struct {
	ID  pgtype.UUID
	No5 string
}

// (scorer role)
func (q *Queries) UpdateStage1try1No5(ctx context.Context, arg UpdateStage1try1No5Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1No5, arg.ID, arg.No5)
	var no5 string
	err := row.Scan(&no5)
	return no5, err
}

const updateStage1try1No6 = `-- name: UpdateStage1try1No6 :one
WITH updated_stage1_results AS (
  UPDATE stage1_results
  SET
    updated_at = NOW()
  WHERE stage1_results.id = $1 
  RETURNING try1_id
) 
UPDATE stage13_tries
SET no6 = $2
WHERE id = (SELECT try1_id FROM stage1_results)
RETURNING no6
`

type UpdateStage1try1No6Params struct {
	ID  pgtype.UUID
	No6 string
}

// (scorer role)
func (q *Queries) UpdateStage1try1No6(ctx context.Context, arg UpdateStage1try1No6Params) (string, error) {
	row := q.db.QueryRow(ctx, updateStage1try1No6, arg.ID, arg.No6)
	var no6 string
	err := row.Scan(&no6)
	return no6, err
}
