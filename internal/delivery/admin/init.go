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
