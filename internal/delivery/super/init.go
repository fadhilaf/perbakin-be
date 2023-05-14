package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"
	superUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"

	"github.com/gin-gonic/gin"
)

type SuperDelivery interface {
	SuperLogin(c *gin.Context)

	MustSuperMiddleware() gin.HandlerFunc

	CheckSuperLogin(c *gin.Context)

	CreateExam(c *gin.Context)
	GetExamsBySuperId(c *gin.Context)
	GetAllExams(c *gin.Context)
	GetExamById(c *gin.Context)
	UpdateExam(c *gin.Context)
	DeleteExam(c *gin.Context)

	GetAllAdmins(c *gin.Context)
	GetAllScorers(c *gin.Context)
	GetAllShooters(c *gin.Context)

	MustExamMiddleware() gin.HandlerFunc

	CreateAdmin(c *gin.Context)
	GetAdminsByExamId(c *gin.Context)
	GetAdminById(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	DeleteAdmin(c *gin.Context)

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

var _ SuperDelivery = &superHandler{}

func NewSuperDelivery(usecase superUsecase.SuperUsecase, adminSuperUsecase adminSuperUsecase.AdminSuperUsecase, sessionUsecase sessionUsecase.SessionUsecase) SuperDelivery {
	return &superHandler{
		Usecase:           usecase,
		AdminSuperUsecase: adminSuperUsecase,
		SessionUsecase:    sessionUsecase,
	}
}

type superHandler struct {
	Usecase           superUsecase.SuperUsecase
	AdminSuperUsecase adminSuperUsecase.AdminSuperUsecase
	SessionUsecase    sessionUsecase.SessionUsecase
}
