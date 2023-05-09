package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ByIdRequest struct {
	ID pgtype.UUID
}
