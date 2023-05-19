package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage0 struct {
	ID        pgtype.UUID `json:"id"`
	ResultID  pgtype.UUID `json:"result_id"`
	Status    string      `json:"status"`
	Series1   []int       `json:"series_1"`
	Series2   []int       `json:"series_2"`
	Series3   []int       `json:"series_3"`
	Series4   []int       `json:"series_4"`
	Series5   []int       `json:"series_5"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type ByResultIdRequest struct {
	ResultID pgtype.UUID `json:"result_id"`
}
