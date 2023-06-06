package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage13Try struct {
	ID         pgtype.UUID `json:"id"`
	Status     string      `json:"status"`
	No1        [][]int     `json:"no_1"`
	No2        [][]int     `json:"no_2"`
	No3        [][]int     `json:"no_3"`
	No4        [][]int     `json:"no_4"`
	No5        [][]int     `json:"no_5"`
	No6        [][]int     `json:"no_6"`
	Checkmarks []bool      `json:"checkmarks"`
}

type Stage1 struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage13Try  `json:"try_1"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage1Full struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage13Try  `json:"try_1"`
	Try2        Stage13Try  `json:"try_2"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage1Relation struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
	Try1ID   pgtype.UUID `json:"try_1_id"`
	Try2ID   pgtype.UUID `json:"try_2_id"`
	IsTry2   bool        `json:"is_try_2"`
}

type Stage123456Try struct {
	Try string `uri:"try" binding:"required,oneof=1 2"`
}
