package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	ID pgtype.UUID `json:"id"`
	User
}

type AdminData struct {
	Admin
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
