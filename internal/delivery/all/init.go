package delivery

import (
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"

	"github.com/gin-gonic/gin"
)

type AllDelivery interface {
	CreateResult(c *gin.Context)

	MustResultMiddleware() gin.HandlerFunc
	GetResultById(c *gin.Context)
}

var _ AllDelivery = &allHandler{}

func NewAllDelivery(usecase allUsecase.AllUsecase) AllDelivery {
	return &allHandler{
		Usecase: usecase,
	}
}

type allHandler struct {
	Usecase allUsecase.AllUsecase
}
