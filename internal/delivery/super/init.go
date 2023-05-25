package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
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

	CreateScorer(c *gin.Context)
	GetScorersByExamId(c *gin.Context)

	GetShootersByExamId(c *gin.Context)

	MustAdminMiddleware() gin.HandlerFunc

	GetAdminById(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	DeleteAdmin(c *gin.Context)

	MustScorerMiddleware() gin.HandlerFunc

	GetScorerById(c *gin.Context)
	UpdateScorer(c *gin.Context)
	DeleteScorer(c *gin.Context)

	GetShootersByScorerId(c *gin.Context)
	CreateShooter(c *gin.Context)

	MustShooterMiddleware() gin.HandlerFunc

	GetShooterById(c *gin.Context)
	UpdateShooter(c *gin.Context)
	UpdateShooterImage(c *gin.Context)
	DeleteShooter(c *gin.Context)
}

var _ SuperDelivery = &superHandler{}

func NewSuperDelivery(usecase superUsecase.SuperUsecase, adminSuperUsecase adminSuperUsecase.AdminSuperUsecase, allUsecase allUsecase.AllUsecase) SuperDelivery {
	return &superHandler{
		Usecase:           usecase,
		AdminSuperUsecase: adminSuperUsecase,
		AllUsecase:        allUsecase,
	}
}

type superHandler struct {
	Usecase           superUsecase.SuperUsecase
	AdminSuperUsecase adminSuperUsecase.AdminSuperUsecase
	AllUsecase        allUsecase.AllUsecase
}
