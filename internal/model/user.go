package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       pgtype.UUID `json:"user_id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
}

type Super struct {
	ID pgtype.UUID `json:"id"`
	User
}

type Admin struct {
	ID pgtype.UUID `json:"id"`
	User
}

type Scorer struct {
	ID pgtype.UUID `json:"id"`
	User
}

type GetByUserIdRequest struct {
	UserID pgtype.UUID `json:"user_id"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
