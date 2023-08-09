package delivery

import (
	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	"github.com/gin-gonic/gin"
)

type AdminDelivery interface {
	AdminLogin(c *gin.Context)

	MustAdminMiddleware() gin.HandlerFunc
}

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase adminUsecase.AdminUsecase) AdminDelivery {
	return &adminHandler{
		Usecase: usecase,
	}
}

type adminHandler struct {
	Usecase adminUsecase.AdminUsecase
}
