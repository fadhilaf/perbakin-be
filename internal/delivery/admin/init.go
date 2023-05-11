package delivery

import (
	adminUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin"
	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"
	"github.com/gin-gonic/gin"
)

type AdminDelivery interface {
	AdminLogin(c *gin.Context)
	MustAdminMiddleware() gin.HandlerFunc
	CheckAdminLogin(c *gin.Context)
}

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase adminUsecase.AdminUsecase, sessionUsecase sessionUsecase.SessionUsecase) AdminDelivery {
	return &adminHandler{
		Usecase:        usecase,
		SessionUsecase: sessionUsecase,
	}
}

type adminHandler struct {
	Usecase        adminUsecase.AdminUsecase
	SessionUsecase sessionUsecase.SessionUsecase
}
