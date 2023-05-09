package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       pgtype.UUID `json:"user_id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
}

type UserDisplayData struct {
	ID        pgtype.UUID `json:"id"`
	Name      string      `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type GetByUserIdRequest struct {
	UserID pgtype.UUID
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,excludes= "`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateUserBodyRequest struct { //ini yg dari http Body ny mending kasi namo kek UpdateUserBodyRequest kek itu
	Username string `json:"username" binding:"required,excludes= "`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	ID   pgtype.UUID
	Body UpdateUserBodyRequest
}
