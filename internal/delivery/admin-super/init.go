package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"

	"github.com/gin-gonic/gin"
)

type AdminSuperDelivery interface {
	MustAdminSuperMiddleware() gin.HandlerFunc

	GetAllScorers(c *gin.Context)

	MustExamMiddleware() gin.HandlerFunc

	CreateScorer(c *gin.Context)
	GetScorersByExamId(c *gin.Context)
	GetScorerById(c *gin.Context)
	UpdateScorer(c *gin.Context)
	DeleteScorer(c *gin.Context)
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
