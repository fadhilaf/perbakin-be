package delivery

import (
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"

	"github.com/gin-gonic/gin"
)

type AllDelivery interface {
	GetShootersByScorerId(c *gin.Context)

	MustShooterMiddleware() gin.HandlerFunc
	GetShooterById(c *gin.Context)

	CreateResult(c *gin.Context)

	MustResultMiddleware() gin.HandlerFunc
	GetResultById(c *gin.Context)

	CreateStage0(c *gin.Context)

	MustStage0Middleware() gin.HandlerFunc
	GetStage0ById(c *gin.Context)
}

var _ AllDelivery = &allHandler{}

func NewAllDelivery(usecase allUsecase.AllUsecase) AllDelivery {
	return &allHandler{
		Usecase: usecase,
	}
}

type allHandler struct {
	Usecase allUsecase.AllUsecase
}
