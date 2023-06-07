package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage13Try struct {
	Status     string  `json:"status"`
	No1        [][]int `json:"no_1"`
	No2        [][]int `json:"no_2"`
	No3        [][]int `json:"no_3"`
	No4        [][]int `json:"no_4"`
	No5        [][]int `json:"no_5"`
	No6        [][]int `json:"no_6"`
	Checkmarks []bool  `json:"checkmarks"`
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

type CreateStage1try2 struct {
	Try2   Stage13Try `json:"try_2"`
	IsTry2 bool       `json:"is_try_2"`
}

type Stage1Relation struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
	IsTry2   bool        `json:"is_try_2"`
}

type Stage123456Try struct {
	Try string `uri:"try" binding:"required,oneof=1 2"`
}

type ByIdAndTryRequest struct {
	ID  pgtype.UUID
	Try string
}

type UpdateStage123456NoUriRequest struct {
	No string `uri:"no" binding:"required,oneof=1 2 3 4 5"`
}

type UpdateStage123NoBodyRequest struct {
	Scores   []int `json:"scores" binding:"required,len=3"`
	Duration []int `json:"duration" binding:"required,len=3,dive,lte=99,gte=0"`
}

type UpdateStage123456NoRequest struct {
	ID     pgtype.UUID `json:"id"`
	Try    string      `json:"try"`
	No     string      `json:"no"`
	Scores string      `json:"scores"`
}

type UpdateStage123456CheckmarksBodyRequest struct {
	Checkmarks []bool `json:"checkmarks" binding:"required,len=6,dive,boolean"`
}

type UpdateStage123456CheckmarksRequest struct {
	ID         pgtype.UUID
	Try        string
	Checkmarks string
}

type UpdateStage123456FinishRequest struct {
	ID          pgtype.UUID
	Try         string
	Success     bool
	ShooterSign pgtype.Text
	ScorerSign  pgtype.Text
}
