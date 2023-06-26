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

	CreateStage1(c *gin.Context)

	MustStage1Middleware() gin.HandlerFunc
	GetStage1ById(c *gin.Context)
	CreateStage1try2(c *gin.Context)

	CreateStage2(c *gin.Context)

	MustStage2Middleware() gin.HandlerFunc
	GetStage2ById(c *gin.Context)
	CreateStage2try2(c *gin.Context)

	CreateStage3(c *gin.Context)

	MustStage3Middleware() gin.HandlerFunc
	GetStage3ById(c *gin.Context)
	CreateStage3try2(c *gin.Context)

	CreateStage4(c *gin.Context)

	MustStage4Middleware() gin.HandlerFunc
	GetStage4ById(c *gin.Context)
	CreateStage4try2(c *gin.Context)
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
