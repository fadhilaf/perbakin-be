package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage5Numbers struct {
	ScoresA  []int `json:"scores_a" binding:"required,len=3"`
	ScoresB  []int `json:"scores_b" binding:"required,len=3"`
	ScoresC  []int `json:"scores_c" binding:"required,len=3"`
	Duration []int `json:"duration" binding:"required,len=3,dive,lte=99,gte=0"`
}

type Stage5Try struct {
	Status     string        `json:"status"`
	No1        Stage5Numbers `json:"no_1"`
	No2        Stage5Numbers `json:"no_2"`
	Checkmarks []bool        `json:"checkmarks"`
}

type Stage5 struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage5Try   `json:"try_1"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage5Full struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage5Try   `json:"try_1"`
	Try2        Stage5Try   `json:"try_2"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
