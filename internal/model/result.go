package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Result struct {
	ID        pgtype.UUID `json:"id"`
	ShooterID pgtype.UUID `json:"shooter_id"`
	Failed    bool        `json:"failed"`
	Stage     string      `json:"stage"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ResultRelation struct {
	ID        pgtype.UUID `json:"id"`
	ShooterID pgtype.UUID `json:"shooter_id"`
}

type ByShooterIdRequest struct {
	ShooterID pgtype.UUID
}

type UpdateResultByShooterIdBodyStringRequest struct {
	Failed string `json:"failed" binding:"required,boolean"`
	Stage  string `json:"stage" binding:"required,number,excludes=.,gte=0,lte=6"`
}

type UpdateResultByShooterIdBodyRequest struct {
	Failed bool
	Stage  string
}

type UpdateResultByShooterIdRequest struct {
	ShooterID pgtype.UUID
	Body      UpdateResultByShooterIdBodyRequest
}
