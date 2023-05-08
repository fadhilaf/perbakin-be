package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	"github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
	MustAdminSuperMiddleware() gin.HandlerFunc
	CreateScorer(c *gin.Context)
	GetScorer(c *gin.Context)
	GetAllScorers(c *gin.Context)
	UpdateScorer(c *gin.Context)
	DeleteScorer(c *gin.Context)
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
