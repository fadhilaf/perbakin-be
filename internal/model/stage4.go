package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage46Numbers struct {
	ScoresA  []int `json:"scores_a" binding:"required,len=3"`
	ScoresB  []int `json:"scores_b" binding:"required,len=3"`
	Duration []int `json:"duration" binding:"required,len=3,dive,lte=99,gte=0"`
}

type Stage46Try struct {
	Status     string         `json:"status"`
	No1        Stage46Numbers `json:"no_1"`
	No2        Stage46Numbers `json:"no_2"`
	No3        Stage46Numbers `json:"no_3"`
	Checkmarks []bool         `json:"checkmarks"`
}

type Stage46 struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage46Try  `json:"try_1"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage46Full struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage46Try  `json:"try_1"`
	Try2        Stage46Try  `json:"try_2"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type CreateStage46try2 struct {
	Try2   Stage46Try `json:"try_2"`
	IsTry2 bool       `json:"is_try_2"`
}

type Stage46UpdateBodyTry struct {
	Status     string         `json:"status" binding:"required,oneof=1 2 3 4"`
	No1        Stage46Numbers `json:"no_1" binding:"required,dive"`
	No2        Stage46Numbers `json:"no_2" binding:"required,dive"`
	No3        Stage46Numbers `json:"no_3" binding:"required,dive"`
	Checkmarks []bool         `json:"checkmarks" binding:"required,len=3,dive,boolean"`
}

type UpdateStage46try1BodyRequest struct {
	Try1 Stage46UpdateBodyTry `json:"try_1" binding:"required,dive"`
}

type UpdateStage46try2BodyRequest struct {
	Try1 Stage46UpdateBodyTry `json:"try_1" binding:"required,dive"`
	Try2 Stage46UpdateBodyTry `json:"try_2" binding:"required,dive"`
}

type UpdateStage46try1Response struct {
	Try1      Stage46Try `json:"try_1"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type UpdateStage46try2Response struct {
	Try1      Stage46Try `json:"try_1"`
	Try2      Stage46Try `json:"try_2"`
	UpdatedAt time.Time  `json:"updated_at"`
}
