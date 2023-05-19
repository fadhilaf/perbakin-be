package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Operator struct {
	ID     pgtype.UUID `json:"id"`
	ExamID pgtype.UUID `json:"exam_id"`
	User
}

type OperatorRelation struct {
	ID     pgtype.UUID `json:"id"`
	ExamID pgtype.UUID `json:"exam_id"`
	UserID pgtype.UUID `json:"user_id"`
}

// untuk GetAllOperators
type OperatorDisplayExamData struct {
	Exam string `json:"exam"`
	Name string `json:"name"`
}

// untuk GetOperatorsByExamId
type OperatorDisplayData struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
}

type OperatorByIdRequest struct {
	ID     pgtype.UUID
	ExamID pgtype.UUID
}

type OperatorBodyRequest struct {
	Username string `json:"username" binding:"required,excludes= "`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type CreateOperatorRequest struct {
	ExamID pgtype.UUID
	Body   OperatorBodyRequest
}

type UpdateOperatorRequest struct {
	ID   pgtype.UUID
	Body OperatorBodyRequest
}
