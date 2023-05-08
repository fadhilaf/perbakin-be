package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Scorer struct {
	ID pgtype.UUID `json:"id"`
	User
}

type ScorerData struct {
  Scorer
  CreatedAt pgtype.Timestamptz `json:"created_at"`
  UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
