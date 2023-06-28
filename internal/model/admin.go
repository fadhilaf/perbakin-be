package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	ID     pgtype.UUID `json:"id"`
	ExamID pgtype.UUID `json:"exam_id"`
	User
}

// Untuk di super admin biar bisa sinkron c.Set("exam") dan c.Set("admin")
type AdminAndExamRelationAndStatus struct {
	ID      pgtype.UUID `json:"id"`
	ExamID  pgtype.UUID `json:"exam_id"`
	UserID  pgtype.UUID `json:"user_id"`
	SuperID pgtype.UUID `json:"super_id"`
	Active  bool        `json:"active"`
}

// get all admins
type AdminDisplayExamData struct {
	Exam string `json:"exam"`
	Name string `json:"name"`
}

// get admins by exam
type AdminDisplayData struct {
	ID   pgtype.UUID `json:"id"`
	Name string      `json:"name"`
}

type CreateAdminBodyRequest struct {
	Username string `json:"username" binding:"required,excludes= "`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type CreateAdminRequest struct {
	ExamID pgtype.UUID
	Body   CreateAdminBodyRequest
}
