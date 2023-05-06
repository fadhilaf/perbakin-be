package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/auth"
	"github.com/gin-gonic/gin"
)

type AuthDelivery interface {
	Logout(c *gin.Context)
}

var _ AuthDelivery = &authHandler{}

func NewAuthDelivery(usecase usecase.AuthUsecase) AuthDelivery {
	return &authHandler{
		Usecase: usecase,
	}
}

type authHandler struct {
	Usecase usecase.AuthUsecase
}
