package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Super struct {
	ID pgtype.UUID `json:"id"`
	User
}
