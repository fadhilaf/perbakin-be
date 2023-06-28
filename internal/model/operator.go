package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type OperatorRelation struct {
	ID     pgtype.UUID `json:"id"`
	ExamID pgtype.UUID `json:"exam_id"`
	UserID pgtype.UUID `json:"user_id"`
}

type OperatorByIdRequest struct {
	ID     pgtype.UUID
	ExamID pgtype.UUID
}

// utk update scorer dan update admin
type OperatorBodyRequest struct {
	Username string `json:"username" binding:"required,excludes= "`
	Password string `json:"password"`
	Name     string `json:"name" binding:"required"`
}

// utk update scorer dan update admin
type UpdateOperatorRequest struct {
	ID       pgtype.UUID
	Username string
	Password pgtype.Text
	Name     string
}
