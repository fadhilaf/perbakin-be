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
	FinishStage0(c *gin.Context)

	UpdateStage1(c *gin.Context)
	UpdateStage1Signs(c *gin.Context)
	DeleteStage1(c *gin.Context)
	DeleteStage1try2(c *gin.Context)
	FinishStage1(c *gin.Context)

	UpdateStage2(c *gin.Context)
	UpdateStage2Signs(c *gin.Context)
	DeleteStage2(c *gin.Context)
	DeleteStage2try2(c *gin.Context)
	FinishStage2(c *gin.Context)

	UpdateStage3(c *gin.Context)
	UpdateStage3Signs(c *gin.Context)
	DeleteStage3(c *gin.Context)
	DeleteStage3try2(c *gin.Context)
	FinishStage3(c *gin.Context)

	UpdateStage4(c *gin.Context)
	UpdateStage4Signs(c *gin.Context)
	DeleteStage4(c *gin.Context)
	DeleteStage4try2(c *gin.Context)
	FinishStage4(c *gin.Context)

	UpdateStage5(c *gin.Context)
	UpdateStage5Signs(c *gin.Context)
	DeleteStage5(c *gin.Context)
	DeleteStage5try2(c *gin.Context)
	FinishStage5(c *gin.Context)

	UpdateStage6(c *gin.Context)
	UpdateStage6Signs(c *gin.Context)
	DeleteStage6(c *gin.Context)
	DeleteStage6try2(c *gin.Context)
	FinishStage6(c *gin.Context)
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
