package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Scorer struct {
	ID        pgtype.UUID `json:"id"`
	ExamID    pgtype.UUID `json:"exam_id"`
	ImagePath string      `json:"image_path"`
	User
}

// Untuk di super admin biar bisa sinkron c.Set("exam") dan c.Set("scorer")
type ScorerRelationAndStatus struct {
	ID     pgtype.UUID `json:"id"`
	ExamID pgtype.UUID `json:"exam_id"`
	UserID pgtype.UUID `json:"user_id"`
	Active bool        `json:"active"`
}

// get all scorers
type ScorerDisplayExamData struct {
	Exam      string `json:"exam"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

// get scorers by exam
type ScorerDisplayData struct {
	ID        pgtype.UUID `json:"id"`
	Name      string      `json:"name"`
	ImagePath string      `json:"image_path"`
}

type CreateScorerBodyRequest struct {
	Username string `form:"username" binding:"required,excludes= "`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
}

type CreateScorerRequest struct {
	ExamID    pgtype.UUID
	ImagePath pgtype.Text
	Body      CreateScorerBodyRequest
}

type UpdateImageRequest struct {
	ID        pgtype.UUID
	ImagePath string
}
