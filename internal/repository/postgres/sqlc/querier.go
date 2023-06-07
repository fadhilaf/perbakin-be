// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// untuk ngebuat admin (super role) TODO: return sebanyak get admin by id
	CreateAdmin(ctx context.Context, arg CreateAdminParams) (CreateAdminRow, error)
	// untuk membuat exam (super role)
	CreateExam(ctx context.Context, arg CreateExamParams) (Exam, error)
	// (all role)
	CreateResult(ctx context.Context, shooterID pgtype.UUID) (Result, error)
	// untuk ngebuat scorer (admin-super role) TODO: return sebanyak get scorer by id
	CreateScorer(ctx context.Context, arg CreateScorerParams) (CreateScorerRow, error)
	// membuat shooter baru (admin-super role)
	CreateShooter(ctx context.Context, arg CreateShooterParams) (Shooter, error)
	// (all role)
	CreateStage0(ctx context.Context, resultID pgtype.UUID) (Stage0Result, error)
	CreateStage1(ctx context.Context, resultID pgtype.UUID) (CreateStage1Row, error)
	CreateStage1try2(ctx context.Context, id pgtype.UUID) (CreateStage1try2Row, error)
	// untuk menghapus exam (super role)
	DeleteExam(ctx context.Context, id pgtype.UUID) error
	// (admin-super role) dibuat by id
	DeleteResult(ctx context.Context, id pgtype.UUID) error
	// untuk menghapus shooter berdasarkan id (admin-super role)
	DeleteShooter(ctx context.Context, id pgtype.UUID) error
	// (admin-super role)
	DeleteStage0(ctx context.Context, id pgtype.UUID) error
	// (admin-super role)
	DeleteStage1(ctx context.Context, id pgtype.UUID) error
	// (admin-super role)
	DeleteStage1try2(ctx context.Context, id pgtype.UUID) error
	// dipake untuk delete user
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	// untuk ngambil data akun admin berdasarkan id (super role)
	GetAdminById(ctx context.Context, id pgtype.UUID) (GetAdminByIdRow, error)
	// untuk ngambil data lengkap admin berdasarkan user id (admin role)
	GetAdminByUserId(ctx context.Context, userID pgtype.UUID) (GetAdminByUserIdRow, error)
	// untuk ngambil data display admin berdasarkan username (admin role)
	GetAdminByUsername(ctx context.Context, username string) (GetAdminByUsernameRow, error)
	// untuk ngambil data relasi admin dan relasi exam berdasarkan user id (all role)
	GetAdminExamRelationByUserId(ctx context.Context, userID pgtype.UUID) (GetAdminExamRelationByUserIdRow, error)
	// untuk ngambil data relasi admin berdasarkan id (all role)
	GetAdminRelationById(ctx context.Context, id pgtype.UUID) (Admin, error)
	// untuk ngambil data akun seluruh admin dalam satu exam (super role)
	GetAdminsByExamId(ctx context.Context, examID pgtype.UUID) ([]GetAdminsByExamIdRow, error)
	// untuk ngambil data display seluruh admin (all role)
	GetAllAdmins(ctx context.Context) ([]GetAllAdminsRow, error)
	// untuk mengambil seluruh exam (super role)
	GetAllExams(ctx context.Context) ([]GetAllExamsRow, error)
	// untuk ngambil data display seluruh scorer (all role)
	GetAllScorers(ctx context.Context) ([]GetAllScorersRow, error)
	// untuk mengambil seluruh shooter (admin-super role)
	GetAllShooters(ctx context.Context) ([]GetAllShootersRow, error)
	// untuk ngambil data display seluruh super admin (all role)
	GetAllSupers(ctx context.Context) ([]GetAllSupersRow, error)
	// untuk mengambil satu data exam (super role)
	GetExamById(ctx context.Context, id pgtype.UUID) (Exam, error)
	// untuk mengambil exam berdasarkan nama untuk cek nama sudah dipakai blum (super role)
	GetExamByName(ctx context.Context, name string) (pgtype.UUID, error)
	// untuk mengambil data relasi exam (all role)
	GetExamRelationById(ctx context.Context, id pgtype.UUID) (GetExamRelationByIdRow, error)
	// untuk mengambil seluruh exam (super role)
	GetExamsBySuperId(ctx context.Context, superID pgtype.UUID) ([]GetExamsBySuperIdRow, error)
	GetResultById(ctx context.Context, id pgtype.UUID) (Result, error)
	GetResultRelationByShooterId(ctx context.Context, shooterID pgtype.UUID) (GetResultRelationByShooterIdRow, error)
	GetResultStatusById(ctx context.Context, id pgtype.UUID) (GetResultStatusByIdRow, error)
	// untuk ngambil data akun scorer berdasarkan id (admin-super role)
	GetScorerById(ctx context.Context, id pgtype.UUID) (GetScorerByIdRow, error)
	// untuk ngambil data lengkap scorer berdasarkan user id (scorer role)
	GetScorerByUserId(ctx context.Context, userID pgtype.UUID) (GetScorerByUserIdRow, error)
	// untuk ngambil data display scorer berdasarkan username (scorer role)
	GetScorerByUsername(ctx context.Context, username string) (GetScorerByUsernameRow, error)
	// untuk ngambil data relasi scorer berdasarkan id (all role)
	GetScorerRelationById(ctx context.Context, id pgtype.UUID) (Scorer, error)
	// untuk ngambil data relasi scorer berdasarkan user id (all role)
	GetScorerRelationByUserId(ctx context.Context, userID pgtype.UUID) (GetScorerRelationByUserIdRow, error)
	// untuk ngambil data akun seluruh scorer dalam satu exam (admin-super role)
	GetScorersByExamId(ctx context.Context, examID pgtype.UUID) ([]GetScorersByExamIdRow, error)
	// untuk mengambil shooter berdasarkan exam_id (admin-super role)
	GetShooterByExamId(ctx context.Context, examID pgtype.UUID) ([]GetShooterByExamIdRow, error)
	// untuk mengambil shooter berdasarkan id (admin-super role)
	GetShooterById(ctx context.Context, id pgtype.UUID) (Shooter, error)
	// untuk mengambil relasi shooter berdasarkan id (all role)
	GetShooterRelationById(ctx context.Context, id pgtype.UUID) (GetShooterRelationByIdRow, error)
	// untuk mengambil shooter berdasarkan scorer_id (all role)
	GetShootersByScorerId(ctx context.Context, scorerID pgtype.UUID) ([]GetShootersByScorerIdRow, error)
	// (all role)
	GetStage0ById(ctx context.Context, id pgtype.UUID) (Stage0Result, error)
	// (all role)
	GetStage0RelationByResultId(ctx context.Context, resultID pgtype.UUID) (GetStage0RelationByResultIdRow, error)
	GetStage0Status(ctx context.Context, id pgtype.UUID) (Stage0Status, error)
	// (all role)
	GetStage1ById(ctx context.Context, id pgtype.UUID) (GetStage1ByIdRow, error)
	// (all role)
	GetStage1RelationByResultId(ctx context.Context, resultID pgtype.UUID) (GetStage1RelationByResultIdRow, error)
	GetStage1try1Status(ctx context.Context, id pgtype.UUID) (Stage13Status, error)
	// (all role)
	GetStage1try2ExistById(ctx context.Context, id pgtype.UUID) (bool, error)
	// (all role)
	GetStage1try2Status(ctx context.Context, id pgtype.UUID) (Stage13Status, error)
	// untuk ngambil data lengkap super admin berdasarkan user id (super role)
	GetSuperByUserId(ctx context.Context, userID pgtype.UUID) (GetSuperByUserIdRow, error)
	// untuk ngambil data display super admin berdasarkan username (super role)
	GetSuperByUsername(ctx context.Context, username string) (GetSuperByUsernameRow, error)
	// untuk ngambil data relasi super admin berdasarkan user id (all role)
	GetSuperRelationByUserId(ctx context.Context, userID pgtype.UUID) (Super, error)
	// dipake untuk mengecek username ketika create user baru
	GetUserByUsername(ctx context.Context, username string) (pgtype.UUID, error)
	// untuk update data akun admin (super role)
	UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (UpdateAdminRow, error)
	// low prio
	UpdateAdminName(ctx context.Context, arg UpdateAdminNameParams) (pgtype.UUID, error)
	// low prio
	UpdateAdminPassword(ctx context.Context, arg UpdateAdminPasswordParams) (pgtype.UUID, error)
	// untuk memperbarui exam (super role)
	UpdateExam(ctx context.Context, arg UpdateExamParams) (UpdateExamRow, error)
	// untuk mengubah status exam (super role)
	UpdateExamStatus(ctx context.Context, id pgtype.UUID) (bool, error)
	// (admin-super role) dibuat by id
	UpdateResult(ctx context.Context, arg UpdateResultParams) (Result, error)
	// (scorer role) dibuat by id, utk update stage
	UpdateResultNextStage(ctx context.Context, arg UpdateResultNextStageParams) error
	// untuk update data akun admin (super role) TODO: return sebanyak get admin by id
	UpdateScorer(ctx context.Context, arg UpdateScorerParams) (UpdateScorerRow, error)
	// low prio
	UpdateScorerName(ctx context.Context, arg UpdateScorerNameParams) (pgtype.UUID, error)
	// low prio
	UpdateScorerPassword(ctx context.Context, arg UpdateScorerPasswordParams) (pgtype.UUID, error)
	// untuk mengupdate shooter berdasarkan id (admin-super role)
	UpdateShooter(ctx context.Context, arg UpdateShooterParams) (UpdateShooterRow, error)
	// untuk mengupdate foto shooter berdasarkan id (admin-super role)
	UpdateShooterImage(ctx context.Context, arg UpdateShooterImageParams) (Shooter, error)
	// (admin-super role)
	UpdateStage0(ctx context.Context, arg UpdateStage0Params) (UpdateStage0Row, error)
	// (scorer role)
	UpdateStage0Checkmarks(ctx context.Context, arg UpdateStage0CheckmarksParams) (string, error)
	// (scorer role)
	UpdateStage0FinishFailed(ctx context.Context, arg UpdateStage0FinishFailedParams) error
	// (scorer role)
	UpdateStage0FinishSuccess(ctx context.Context, arg UpdateStage0FinishSuccessParams) error
	// (scorer role)
	UpdateStage0NextSeries(ctx context.Context, arg UpdateStage0NextSeriesParams) error
	// (scorer role)
	UpdateStage0Series1(ctx context.Context, arg UpdateStage0Series1Params) (string, error)
	// (scorer role)
	UpdateStage0Series2(ctx context.Context, arg UpdateStage0Series2Params) (string, error)
	// (scorer role)
	UpdateStage0Series3(ctx context.Context, arg UpdateStage0Series3Params) (string, error)
	// (scorer role)
	UpdateStage0Series4(ctx context.Context, arg UpdateStage0Series4Params) (string, error)
	// (scorer role)
	UpdateStage0Series5(ctx context.Context, arg UpdateStage0Series5Params) (string, error)
	// (admin-super role)
	UpdateStage0Signs(ctx context.Context, arg UpdateStage0SignsParams) (UpdateStage0SignsRow, error)
	// (admin-super role)
	UpdateStage1(ctx context.Context, arg UpdateStage1Params) (UpdateStage1Row, error)
	// (scorer role)
	UpdateStage1NextTry(ctx context.Context, id pgtype.UUID) error
	// (admin-super role)
	UpdateStage1Signs(ctx context.Context, arg UpdateStage1SignsParams) (UpdateStage1SignsRow, error)
	// (scorer role)
	UpdateStage1try1Checkmarks(ctx context.Context, arg UpdateStage1try1CheckmarksParams) (string, error)
	// (scorer role)
	UpdateStage1try1FinishFailed(ctx context.Context, arg UpdateStage1try1FinishFailedParams) error
	// (scorer role)
	UpdateStage1try1FinishSuccess(ctx context.Context, arg UpdateStage1try1FinishSuccessParams) error
	// (scorer role)
	UpdateStage1try1NextNo(ctx context.Context, arg UpdateStage1try1NextNoParams) error
	// (scorer role)
	UpdateStage1try1No1(ctx context.Context, arg UpdateStage1try1No1Params) (string, error)
	// (scorer role)
	UpdateStage1try1No2(ctx context.Context, arg UpdateStage1try1No2Params) (string, error)
	// (scorer role)
	UpdateStage1try1No3(ctx context.Context, arg UpdateStage1try1No3Params) (string, error)
	// (scorer role)
	UpdateStage1try1No4(ctx context.Context, arg UpdateStage1try1No4Params) (string, error)
	// (scorer role)
	UpdateStage1try1No5(ctx context.Context, arg UpdateStage1try1No5Params) (string, error)
	// (scorer role)
	UpdateStage1try1No6(ctx context.Context, arg UpdateStage1try1No6Params) (string, error)
	// (scorer role)
	UpdateStage1try2Checkmarks(ctx context.Context, arg UpdateStage1try2CheckmarksParams) (string, error)
	// (scorer role)
	UpdateStage1try2FinishFailed(ctx context.Context, arg UpdateStage1try2FinishFailedParams) error
	// (scorer role)
	UpdateStage1try2FinishSuccess(ctx context.Context, arg UpdateStage1try2FinishSuccessParams) error
	// (scorer role)
	UpdateStage1try2NextNo(ctx context.Context, arg UpdateStage1try2NextNoParams) error
	// (scorer role)
	UpdateStage1try2No1(ctx context.Context, arg UpdateStage1try2No1Params) (string, error)
	// (scorer role)
	UpdateStage1try2No2(ctx context.Context, arg UpdateStage1try2No2Params) (string, error)
	// (scorer role)
	UpdateStage1try2No3(ctx context.Context, arg UpdateStage1try2No3Params) (string, error)
	// (scorer role)
	UpdateStage1try2No4(ctx context.Context, arg UpdateStage1try2No4Params) (string, error)
	// (scorer role)
	UpdateStage1try2No5(ctx context.Context, arg UpdateStage1try2No5Params) (string, error)
	// (scorer role)
	UpdateStage1try2No6(ctx context.Context, arg UpdateStage1try2No6Params) (string, error)
}

var _ Querier = (*Queries)(nil)
