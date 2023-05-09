package delivery

import (
	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
	"github.com/gin-gonic/gin"
)

type AdminDelivery interface {
	AdminLogin(c *gin.Context)
	MustAdminMiddleware() gin.HandlerFunc
	CheckAdminLogin(c *gin.Context)
}

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase adminUsecase.AdminUsecase, allUsecase allUsecase.AllUsecase) AdminDelivery {
	return &adminHandler{
		Usecase:    usecase,
		AllUsecase: allUsecase,
	}
}

type adminHandler struct {
	Usecase    adminUsecase.AdminUsecase
	AllUsecase allUsecase.AllUsecase
}
