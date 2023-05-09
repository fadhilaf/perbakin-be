package model

type LoginRequest struct {
	Username string `json:"username" binding:"required,excludes= "`
	Password string `json:"password" binding:"required"`
}
