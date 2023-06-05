package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage1 struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
	Status   string      `json:"status"`
	// Scores1     []int       `json:"scores_1"`
	// Duration1   []int       `json:"duration_1"`
	// Scores2     []int       `json:"scores_2"`
	// Duration2   []int       `json:"duration_2"`
	// Scores3     []int       `json:"scores_3"`
	// Duration3   []int       `json:"duration_3"`
	// Scores4     []int       `json:"scores_4"`
	// Duration4   []int       `json:"duration_4"`
	// Scores5     []int       `json:"scores_5"`
	// Duration5   []int       `json:"duration_5"`
	// Scores6     []int       `json:"scores_6"`
	// Duration6   []int       `json:"duration_6"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
