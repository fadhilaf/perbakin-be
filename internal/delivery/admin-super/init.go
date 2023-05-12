package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"

	"github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
	MustAdminSuperMiddleware() gin.HandlerFunc

	GetAllScorers(c *gin.Context)
	GetAllShooters(c *gin.Context)

	MustExamMiddleware() gin.HandlerFunc

	CreateScorer(c *gin.Context)
	GetScorersByExamId(c *gin.Context)
	GetScorerById(c *gin.Context)
	UpdateScorer(c *gin.Context)
	DeleteScorer(c *gin.Context)

	MustScorerMiddleware() gin.HandlerFunc

	CreateShooter(c *gin.Context)
	GetShooterById(c *gin.Context)
	GetShootersByExamId(c *gin.Context)
	UpdateShooter(c *gin.Context)
	DeleteShooter(c *gin.Context)
}

var _ AdminSuperDelivery = &adminSuperHandler{}

func NewAdminSuperDelivery(usecase adminSuperUsecase.AdminSuperUsecase, sessionUsecase sessionUsecase.SessionUsecase) AdminSuperDelivery {
	return &adminSuperHandler{
		Usecase:        usecase,
		SessionUsecase: sessionUsecase,
	}
}

type adminSuperHandler struct {
	Usecase        adminSuperUsecase.AdminSuperUsecase
	SessionUsecase sessionUsecase.SessionUsecase
}
