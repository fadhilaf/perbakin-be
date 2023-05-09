package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       pgtype.UUID `json:"user_id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
}

type UserByUserIdRequest struct {
	UserID pgtype.UUID `json:"user_id"`
}
