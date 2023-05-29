package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"

	"github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
	GetAllScorers(c *gin.Context)
	GetAllShooters(c *gin.Context)

	GetScorersByExamId(c *gin.Context)
	GetShootersByExamId(c *gin.Context)

	CreateScorer(c *gin.Context)

	MustScorerMiddleware() gin.HandlerFunc
	GetScorerById(c *gin.Context)
	UpdateScorer(c *gin.Context)
	DeleteScorer(c *gin.Context)

	CreateShooter(c *gin.Context)

	UpdateShooter(c *gin.Context)
	UpdateShooterImage(c *gin.Context)
	DeleteShooter(c *gin.Context)

	UpdateResult(c *gin.Context)
	DeleteResult(c *gin.Context)

	UpdateStage0(c *gin.Context)
	UpdateStage0Signs(c *gin.Context)
	DeleteStage0(c *gin.Context)
}

var _ AdminSuperDelivery = &adminSuperHandler{}

func NewAdminSuperDelivery(usecase adminSuperUsecase.AdminSuperUsecase) AdminSuperDelivery {
	return &adminSuperHandler{
		Usecase: usecase,
	}
}

type adminSuperHandler struct {
	Usecase adminSuperUsecase.AdminSuperUsecase
}
