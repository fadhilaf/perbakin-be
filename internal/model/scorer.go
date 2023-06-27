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
