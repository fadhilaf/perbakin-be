package delivery

import (
	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"
	"github.com/gin-gonic/gin"
)

type AdminDelivery interface {
	AdminLogin(c *gin.Context)

	MustAdminMiddleware() gin.HandlerFunc

	CheckAdminLogin(c *gin.Context)

	GetAllScorers(c *gin.Context)
	GetAllShooters(c *gin.Context)

	CreateScorer(c *gin.Context)
	GetScorersByExamId(c *gin.Context)
	GetScorerById(c *gin.Context)
	UpdateScorer(c *gin.Context)
	DeleteScorer(c *gin.Context)

	GetShootersByExamId(c *gin.Context)

	MustScorerMiddleware() gin.HandlerFunc

	CreateShooter(c *gin.Context)
	GetShooterById(c *gin.Context)
	UpdateShooter(c *gin.Context)
	DeleteShooter(c *gin.Context)
}

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase adminUsecase.AdminUsecase, adminSuperUsecase adminSuperUsecase.AdminSuperUsecase, sessionUsecase sessionUsecase.SessionUsecase) AdminDelivery {
	return &adminHandler{
		Usecase:           usecase,
		AdminSuperUsecase: adminSuperUsecase,
		SessionUsecase:    sessionUsecase,
	}
}

type adminHandler struct {
	Usecase           adminUsecase.AdminUsecase
	AdminSuperUsecase adminSuperUsecase.AdminSuperUsecase
	SessionUsecase    sessionUsecase.SessionUsecase
}
