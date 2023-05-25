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

type Stage0Relation struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
}

type ByResultIdRequest struct {
	ResultID pgtype.UUID
}

type UpdateStage0SeriesBodyRequest struct {
	Series string `json:"series" binding:"required,oneof=1 2 3 4 5"`
	Scores []int  `json:"scores" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
}

type UpdateStage0SeriesRequest struct {
	ID     pgtype.UUID
	Series string
	Scores string
}
