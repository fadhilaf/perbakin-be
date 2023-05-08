package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	"github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
	GetAllAdmins(c *gin.Context)
	MustAdminSuperMiddleware() gin.HandlerFunc
}

var _ AdminSuperDelivery = &adminSuperHandler{}

func NewAdminSuperDelivery(usecase usecase.AdminSuperUsecase) AdminSuperDelivery {
	return &adminSuperHandler{
		Usecase: usecase,
	}
}

type adminSuperHandler struct {
	Usecase usecase.AdminSuperUsecase
}
