package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Shooter struct {
	ID        pgtype.UUID `json:"id"`
	ScorerID  pgtype.UUID `json:"scorer_id"`
	Name      string      `json:"name"`
	ImagePath string      `json:"image_path"`
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
	Exam      string `json:"exam"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
	Province  string `json:"province"`
	Club      string `json:"club"`
}

// untuk GetShootersByExamId
type ShooterDisplayScorerData struct {
	ID        pgtype.UUID `json:"id"`
	ScorerID  pgtype.UUID `json:"scorer_id"`
	Scorer    string      `json:"scorer"`
	Name      string      `json:"name"`
	ImagePath string      `json:"image_path"`
	Province  string      `json:"province"`
	Club      string      `json:"club"`
}

// untuk GetShootersByScorerId
type ShooterDisplayData struct {
	ID        pgtype.UUID `json:"id"`
	Name      string      `json:"name"`
	ImagePath string      `json:"image_path"`
	Province  string      `json:"province"`
	Club      string      `json:"club"`
}

type ByScorerIdRequest struct {
	ScorerID pgtype.UUID
}

type ShooterByIdRequest struct {
	ID       pgtype.UUID
	ScorerID pgtype.UUID
}

type CreateShooterBodyRequest struct {
	Name     string `form:"name" binding:"required"`
	Province string `form:"province" binding:"required"`
	Club     string `form:"club" binding:"required"`
}

type CreateShooterRequest struct {
	ScorerID  pgtype.UUID
	ImagePath pgtype.Text
	Body      CreateShooterBodyRequest
}

type UpdateShooterBodyRequest struct {
	ScorerID pgtype.UUID `json:"scorer_id" binding:"required,uuid"`
	Name     string      `json:"name" binding:"required"`
	Province string      `json:"province" binding:"required"`
	Club     string      `json:"club" binding:"required"`
}

type UpdateShooterImageRequest struct {
	ID        pgtype.UUID
	ImagePath string
}

type UpdateShooterRequest struct {
	ID     pgtype.UUID
	ExamID pgtype.UUID
	Body   UpdateShooterBodyRequest
}
