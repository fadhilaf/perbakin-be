package delivery

import (
	usecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"
	"github.com/gin-gonic/gin"
)

type SuperDelivery interface {
	SuperLogin(c *gin.Context)
	MustSuperMiddleware() gin.HandlerFunc
	CheckSuperLogin(c *gin.Context)

	CreateAdmin(c *gin.Context)
	GetAdminById(c *gin.Context)
	GetAllAdmins(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	DeleteAdmin(c *gin.Context)

	GetExamById(c *gin.Context)
	CreateExam(c *gin.Context)
	UpdateExam(c *gin.Context)
	DeleteExam(c *gin.Context)
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
