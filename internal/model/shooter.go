package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Shooter struct {
	ID        pgtype.UUID `json:"id"`
	ScorerID  pgtype.UUID `json:"scorer_id"`
	Name      string      `json:"name"`
	Province  string      `json:"province"`
	Club      string      `json:"club"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ShooterRelation struct {
	ID       pgtype.UUID `json:"id"`
	ScorerID pgtype.UUID `json:"scorer_id"`
}

// untuk GetAllShooters
type ShooterDisplayExamData struct {
	Exam     string `json:"exam"`
	Name     string `json:"name"`
	Province string `json:"province"`
	Club     string `json:"club"`
}

// untuk GetShootersByExamId
type ShooterDisplayScorerData struct {
	ID       pgtype.UUID `json:"id"`
	ScorerID pgtype.UUID `json:"scorer_id"`
	Scorer   string      `json:"scorer"`
	Name     string      `json:"name"`
	Province string      `json:"province"`
	Club     string      `json:"club"`
}

// untuk GetShootersByScorerId
type ShooterDisplayData struct {
	ID       pgtype.UUID `json:"id"`
	Name     string      `json:"name"`
	Province string      `json:"province"`
	Club     string      `json:"club"`
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
	ID     pgtype.UUID
	ExamID pgtype.UUID
	Body   UpdateShooterBodyRequest
}
