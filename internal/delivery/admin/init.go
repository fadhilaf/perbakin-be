package delivery

import (
	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
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

	GetShootersByExamId(c *gin.Context)

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

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase adminUsecase.AdminUsecase, adminSuperUsecase adminSuperUsecase.AdminSuperUsecase, allUsecase allUsecase.AllUsecase) AdminDelivery {
	return &adminHandler{
		Usecase:           usecase,
		AdminSuperUsecase: adminSuperUsecase,
		AllUsecase:        allUsecase,
	}
}

type adminHandler struct {
	Usecase           adminUsecase.AdminUsecase
	AdminSuperUsecase adminSuperUsecase.AdminSuperUsecase
	AllUsecase        allUsecase.AllUsecase
}
