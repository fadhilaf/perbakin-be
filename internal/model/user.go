package model

import (
	"github.com/jackc/pgx/v5/pgtype"

	"time"
)

type User struct {
	ID        pgtype.UUID `json:"user_id"`
	Username  string      `json:"username"`
	Name      string      `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type UserByUserIdRequest struct {
	UserID pgtype.UUID `json:"user_id"`
}
