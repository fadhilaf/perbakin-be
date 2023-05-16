package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Exam struct {
	ID        pgtype.UUID `json:"id"`
	SuperID   pgtype.UUID `json:"super_id"`
	Name      string      `json:"name"`
	Location  string      `json:"location"`
	Organizer string      `json:"organizer"`
	Begin     pgtype.Date `json:"begin"`
	Finish    pgtype.Date `json:"finish"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ExamDisplaySuperData struct {
	Name      string      `json:"name"`
	Super     string      `json:"super"`
	Location  string      `json:"location"`
	Organizer string      `json:"organizer"`
	Begin     pgtype.Date `json:"begin"`
	Finish    pgtype.Date `json:"finish"`
}

type ExamDisplayData struct {
	ID        pgtype.UUID `json:"id"`
	Name      string      `json:"name"`
	Location  string      `json:"location"`
	Organizer string      `json:"organizer"`
	Begin     pgtype.Date `json:"begin"`
	Finish    pgtype.Date `json:"finish"`
}

type ExamRelation struct {
	ID      pgtype.UUID `json:"id"`
	SuperID pgtype.UUID `json:"super_id"`
}

// Bentuk struct yang biso di validasi itu type ny harus string
type CreateExamBodyStringRequest struct {
	Name      string `json:"name" binding:"required"`
	Location  string `json:"location" binding:"required"`
	Organizer string `json:"organizer" binding:"required"`
	Begin     string `json:"begin" binding:"required,datetime=2006-01-02"` //validasi format date ny niru time.Time golang
	Finish    string `json:"finish" binding:"required,datetime=2006-01-02"`
}

type CreateExamBodyRequest struct {
	Name      string
	Location  string
	Organizer string
	Begin     pgtype.Date
	Finish    pgtype.Date
}

type CreateExamRequest struct {
	SuperID pgtype.UUID
	Body    CreateExamBodyRequest
}

type GetExamsBySuperIdRequest struct {
	SuperID pgtype.UUID
}

// Bentuk struct yang biso di validasi itu type ny harus string
type UpdateExamBodyStringRequest struct {
	Name      string `json:"name" binding:"required"`
	Location  string `json:"location" binding:"required"`
	Organizer string `json:"organizer" binding:"required"`
	Begin     string `json:"begin" binding:"required,datetime=2006-01-02"`
	Finish    string `json:"finish" binding:"required,datetime=2006-01-02"`
}

type UpdateExamBodyRequest struct {
	Name      string
	Location  string
	Organizer string
	Begin     pgtype.Date
	Finish    pgtype.Date
}

type UpdateExamRequest struct {
	ID      pgtype.UUID
	SuperID pgtype.UUID
	Body    UpdateExamBodyRequest
}

type DeleteExamRequest struct {
	ID      pgtype.UUID
	SuperID pgtype.UUID
}
