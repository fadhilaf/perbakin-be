package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage0 struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Status      string      `json:"status"`
	Series1     []int       `json:"series_1"`
	Series2     []int       `json:"series_2"`
	Series3     []int       `json:"series_3"`
	Series4     []int       `json:"series_4"`
	Series5     []int       `json:"series_5"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage0Relation struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
}

type FinishStage0Response struct {
	Status      string      `json:"status"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
}

type ByResultIdRequest struct {
	ResultID pgtype.UUID
}

type UpdateStage0SeriesUriRequest struct {
	Series string `uri:"series" binding:"required,oneof=1 2 3 4 5"`
}

type UpdateStage0SeriesBodyRequest struct {
	Scores []int `json:"scores" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
}

type UpdateStage0SeriesRequest struct {
	ID     pgtype.UUID
	Series string
	Scores string
}

type UpdateStage0FinishBodyRequest struct {
	Success bool `form:"success" binding:"required,boolean"`
}

type UpdateStage0FinishRequest struct {
	ID          pgtype.UUID
	ResultID    pgtype.UUID
	Body        UpdateStage0FinishBodyRequest
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}
