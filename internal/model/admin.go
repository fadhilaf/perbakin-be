package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	ID     pgtype.UUID `json:"id"`
	ExamID pgtype.UUID `json:"exam_id"`
	User
}
