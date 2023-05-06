package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"
	"github.com/gin-gonic/gin"
)

type SuperDelivery interface {
	SuperLogin(c *gin.Context)
	CheckSuperLogin(c *gin.Context)
}

var _ SuperDelivery = &superHandler{}

func NewSuperDelivery(usecase usecase.SuperUsecase) SuperDelivery {
	return &superHandler{
		Usecase: usecase,
	}
}

type superHandler struct {
	Usecase usecase.SuperUsecase
}
