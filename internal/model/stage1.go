package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage123Numbers struct {
	Scores   []int `json:"scores" binding:"required,len=3"`
	Duration []int `json:"duration" binding:"required,len=3,dive,lte=99,gte=0"`
}

type Stage13Try struct {
	Status     string          `json:"status"`
	No1        Stage123Numbers `json:"no_1"`
	No2        Stage123Numbers `json:"no_2"`
	No3        Stage123Numbers `json:"no_3"`
	No4        Stage123Numbers `json:"no_4"`
	No5        Stage123Numbers `json:"no_5"`
	No6        Stage123Numbers `json:"no_6"`
	Checkmarks []bool          `json:"checkmarks"`
}

type Stage13 struct {
	ID          pgtype.UUID `json:"id"`
	ResultID    pgtype.UUID `json:"result_id"`
	Try1        Stage13Try  `json:"try_1"`
	IsTry2      bool        `json:"is_try_2"`
	ShooterSign pgtype.Text `json:"shooter_sign"`
	ScorerSign  pgtype.Text `json:"scorer_sign"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Stage13Full struct {
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

type CreateStage13try2 struct {
	Try2   Stage13Try `json:"try_2"`
	IsTry2 bool       `json:"is_try_2"`
}

type Stage123456Relation struct {
	ID       pgtype.UUID `json:"id"`
	ResultID pgtype.UUID `json:"result_id"`
	IsTry2   bool        `json:"is_try_2"`
}

type Stage13UpdateBodyTry struct {
	Status     string          `json:"status" binding:"required,oneof=1 2 3 4 5 6"`
	No1        Stage123Numbers `json:"no_1" binding:"required,dive"`
	No2        Stage123Numbers `json:"no_2" binding:"required,dive"`
	No3        Stage123Numbers `json:"no_3" binding:"required,dive"`
	No4        Stage123Numbers `json:"no_4" binding:"required,dive"`
	No5        Stage123Numbers `json:"no_5" binding:"required,dive"`
	No6        Stage123Numbers `json:"no_6" binding:"required,dive"`
	Checkmarks []bool          `json:"checkmarks" binding:"required,len=6,dive,boolean"`
}

type UpdateStage13try1BodyRequest struct {
	Try1 Stage13UpdateBodyTry `json:"try_1" binding:"required,dive"`
}

type UpdateStage13try2BodyRequest struct {
	Try1 Stage13UpdateBodyTry `json:"try_1" binding:"required,dive"`
	Try2 Stage13UpdateBodyTry `json:"try_2" binding:"required,dive"`
}

type Stage13TryString struct {
	Status     string `json:"status"`
	No1        string `json:"no_1"`
	No2        string `json:"no_2"`
	No3        string `json:"no_3"`
	No4        string `json:"no_4"`
	No5        string `json:"no_5"`
	No6        string `json:"no_6"`
	Checkmarks string `json:"checkmarks"`
}

type UpdateStage13try1Request struct {
	ID   pgtype.UUID
	Try1 Stage13TryString
}

type UpdateStage13try2Request struct {
	ID   pgtype.UUID
	Try1 Stage13TryString
	Try2 Stage13TryString
}

type UpdateStage13Response struct {
	Try1      Stage13Try `json:"try_1"`
	Try2      Stage13Try `json:"try_2"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Stage123456TryRequestParam struct {
	Try string `uri:"try" binding:"required,oneof=1 2"`
}

type ByIdAndTryRequest struct {
	ID  pgtype.UUID
	Try string
}

type UpdateStage13NoUriRequest struct {
	No string `uri:"no" binding:"required,oneof=1 2 3 4 5 6"`
}

// scores in this struct means (( score_a, score_b, score_c ), checkmarks)
type UpdateStage123456NoRequest struct {
	ID                pgtype.UUID
	Try               string
	No                string
	ScoresAndDuration string
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
