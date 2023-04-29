// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateAdmin(ctx context.Context, arg CreateAdminParams) (pgconn.CommandTag, error)
	CreateScorer(ctx context.Context, arg CreateScorerParams) (pgconn.CommandTag, error)
	DeleteAdmin(ctx context.Context, userID pgtype.UUID) error
	DeleteScorer(ctx context.Context, userID pgtype.UUID) error
	GetAdminByUserId(ctx context.Context, userID pgtype.UUID) (GetAdminByUserIdRow, error)
	GetAdminByUsername(ctx context.Context, username string) (GetAdminByUsernameRow, error)
	GetAdmins(ctx context.Context) ([]GetAdminsRow, error)
	GetScorer(ctx context.Context, userID pgtype.UUID) (Scorer, error)
	GetScorerByUsername(ctx context.Context, username string) (GetScorerByUsernameRow, error)
	GetScorers(ctx context.Context) ([]Scorer, error)
	GetSuperByUsername(ctx context.Context, username string) (GetSuperByUsernameRow, error)
}

var _ Querier = (*Queries)(nil)
