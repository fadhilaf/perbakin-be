package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"

	"github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
	MustResultMiddleware() gin.HandlerFunc
	UpdateResult(c *gin.Context)
	DeleteResult(c *gin.Context)

	MustStage0Middleware() gin.HandlerFunc
	UpdateStage0(c *gin.Context)
	UpdateStage0Signs(c *gin.Context)
	DeleteStage0(c *gin.Context)
}

var _ AdminSuperDelivery = &adminSuperHandler{}

func NewAdminSuperDelivery(usecase adminSuperUsecase.AdminSuperUsecase, allUsecase allUsecase.AllUsecase) AdminSuperDelivery {
	return &adminSuperHandler{
		Usecase:    usecase,
		AllUsecase: allUsecase,
	}
}

type adminSuperHandler struct {
	Usecase    adminSuperUsecase.AdminSuperUsecase
	AllUsecase allUsecase.AllUsecase
}
