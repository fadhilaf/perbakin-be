package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Stage5Numbers struct {
	ScoresA  []int `json:"scores_a" binding:"required,len=3,dive,oneof=0 1 2"`
	ScoresB  []int `json:"scores_b" binding:"required,len=3,dive,oneof=0 1 2"`
	ScoresC  []int `json:"scores_c" binding:"required,len=3,dive,oneof=0 1 2"`
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

type CreateStage5try2 struct {
	Try2   Stage5Try `json:"try_2"`
	IsTry2 bool      `json:"is_try_2"`
}

type UpdateStage5NoUriRequest struct {
	No string `uri:"no" binding:"required,oneof=1 2"`
}

type UpdateStage5CheckmarksBodyRequest struct {
	Checkmarks []bool `json:"checkmarks" binding:"required,len=2,dive,boolean"`
}

type Stage5TryString struct {
	Status     string `json:"status"`
	No1        string `json:"no_1"`
	No2        string `json:"no_2"`
	Checkmarks string `json:"checkmarks"`
}

type UpdateStage5try1Request struct {
	ID   pgtype.UUID
	Try1 Stage5TryString
}

type UpdateStage5try2Request struct {
	ID   pgtype.UUID
	Try1 Stage5TryString
	Try2 Stage5TryString
}

type Stage5UpdateBodyTry struct {
	Status     string        `json:"status" binding:"required,oneof=1 2 3"`
	No1        Stage5Numbers `json:"no_1" binding:"required,dive"`
	No2        Stage5Numbers `json:"no_2" binding:"required,dive"`
	Checkmarks []bool        `json:"checkmarks" binding:"required,len=2,dive,boolean"`
}

type UpdateStage5try1BodyRequest struct {
	Try1 Stage5UpdateBodyTry `json:"try_1" binding:"required,dive"`
}

type UpdateStage5try2BodyRequest struct {
	Try1 Stage5UpdateBodyTry `json:"try_1" binding:"required,dive"`
	Try2 Stage5UpdateBodyTry `json:"try_2" binding:"required,dive"`
}

type UpdateStage5try1Response struct {
	Try1      Stage5Try `json:"try_1"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateStage5try2Response struct {
	Try1      Stage5Try `json:"try_1"`
	Try2      Stage5Try `json:"try_2"`
	UpdatedAt time.Time `json:"updated_at"`
}
