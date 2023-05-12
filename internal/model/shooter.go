package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Shooter struct {
	ID       pgtype.UUID `json:"user_id"`
	ScorerID pgtype.UUID `json:"scorer_id"`
	Name     string      `json:"name"`
	Province string      `json:"province"`
	Club     string      `json:"club"`
}

type ShooterDisplayData struct {
	ID        pgtype.UUID `json:"user_id"`
	Exam      string      `json:"exam"`
	Name      string      `json:"name"`
	Province  string      `json:"province"`
	Club      string      `json:"club"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ByScorerIdRequest struct {
	ScorerID pgtype.UUID
}

type ShooterByIdRequest struct {
	ID       pgtype.UUID
	ScorerID pgtype.UUID
}

type CreateShooterBodyRequest struct {
	Name     string `json:"name"`
	Province string `json:"province"`
	Club     string `json:"club"`
}

type CreateShooterRequest struct {
	ScorerID pgtype.UUID
	Body     CreateShooterBodyRequest
}

type UpdateShooterBodyRequest struct {
	ScorerID pgtype.UUID `json:"scorer_id"`
	Name     string      `json:"name"`
	Province string      `json:"province"`
	Club     string      `json:"club"`
}

type UpdateShooterRequest struct {
	ID       pgtype.UUID
	ExamID   pgtype.UUID
	ScorerID pgtype.UUID
	Body     UpdateShooterBodyRequest
}
