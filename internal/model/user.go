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

type GetUserById struct {
	ID pgtype.UUID
}

type GetByUserIdRequest struct {
	UserID pgtype.UUID `json:"user_id"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UpdateUserDataRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UpdateUserRequest struct {
	ID   pgtype.UUID
	Data UpdateUserDataRequest
}

type DeleteUserRequest struct {
	ID pgtype.UUID
}
