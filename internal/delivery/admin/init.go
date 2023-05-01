package delivery

import (
	usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/admin"
	"github.com/gin-gonic/gin"
)

type AdminDelivery interface {
	AdminLogin(c *gin.Context)
}

var _ AdminDelivery = &adminHandler{}

func NewAdminDelivery(usecase usecase.AdminUsecase) AdminDelivery {
	return &adminHandler{
		usecase: usecase,
	}
}

type adminHandler struct {
	usecase usecase.AdminUsecase
}
