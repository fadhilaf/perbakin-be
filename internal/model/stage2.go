package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage2Try struct {
	Status     string          `json:"status"`
	No1        Stage123Numbers `json:"no_1"`
	No2        Stage123Numbers `json:"no_2"`
	No3        Stage123Numbers `json:"no_3"`
	Checkmarks []bool          `json:"checkmarks"`
}

type Stage2 struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage2Try   `json:"try_1"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage2Full struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage2Try   `json:"try_1"`
	Try2        Stage2Try   `json:"try_2"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type CreateStage2try2 struct {
	Try2   Stage2Try `json:"try_2"`
	IsTry2 bool      `json:"is_try_2"`
}

type UpdateStage246NoUriRequest struct {
	No string `uri:"no" binding:"required,oneof=1 2 3"`
}

type UpdateStage246CheckmarksBodyRequest struct {
	Checkmarks []bool `json:"checkmarks" binding:"required,len=3,dive,boolean"`
}
