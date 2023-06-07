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
	Checkmarks  []bool      `json:"checkmarks"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage0Relation struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
}

type ByResultIdRequest struct {
	ResultID pgtype.UUID
}

type UpdateStage0BodyRequest struct {
	Status     string `json:"status" binding:"required,oneof=1 2 3 4 5 6"`
	Series1    []int  `json:"series_1" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
	Series2    []int  `json:"series_2" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
	Series3    []int  `json:"series_3" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
	Series4    []int  `json:"series_4" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
	Series5    []int  `json:"series_5" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
	Checkmarks []bool `json:"checkmarks" binding:"required,dive,boolean"`
}

type UpdateStage0Request struct {
	ID         pgtype.UUID
	Status     string
	Series1    string
	Series2    string
	Series3    string
	Series4    string
	Series5    string
	Checkmarks string
}

type UpdateStage0Response struct {
	ID         pgtype.UUID `json:"id"`
	ResultID   pgtype.UUID `json:"result_id"`
	Status     string      `json:"status"`
	Series1    []int       `json:"series_1"`
	Series2    []int       `json:"series_2"`
	Series3    []int       `json:"series_3"`
	Series4    []int       `json:"series_4"`
	Series5    []int       `json:"series_5"`
	Checkmarks []bool      `json:"checkmarks"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type UpdateStageSignsResponse struct {
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type UpdateStageSignsRequest struct {
	ID          pgtype.UUID
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}

type UpdateStage0SeriesUriRequest struct {
	Series string `uri:"series" binding:"required,oneof=1 2 3 4 5"`
}

type UpdateStageScoresBodyRequest struct {
	Scores []int `json:"scores" binding:"required,dive,oneof=0 1 2 3 4 5 6 7 8 9 10"`
}

type UpdateStage0SeriesRequest struct {
	ID     pgtype.UUID
	Series string
	Scores string
}

// yang ini kalo dio salah type, error message nyo jelek, tapi salah dewek lh
type UpdateStageCheckmarksBodyRequest struct {
	Checkmarks []bool `json:"checkmarks" binding:"required,dive,boolean"`
}

type UpdateStage0CheckmarksRequest struct {
	ID         pgtype.UUID
	Checkmarks string
}

// samo, kalo pake bool kalo kito kasih value 'false' (boolean json) dio malah jadi kosong dianggapny. jadi pake *bool
type UpdateStageFinishBodyRequest struct {
	Success *bool `form:"success" binding:"required,boolean"`
}

type UpdateStageFinishRequest struct {
	ID          pgtype.UUID
	Success     bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}
