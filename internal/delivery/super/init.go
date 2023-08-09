package delivery

import (
	superUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"

	"github.com/gin-gonic/gin"
)

type SuperDelivery interface {
	SuperLogin(c *gin.Context)

	MustSuperMiddleware() gin.HandlerFunc

	CheckSuperLogin(c *gin.Context)

	GetAllExams(c *gin.Context)

	GetExamsBySuperId(c *gin.Context)
	CreateExam(c *gin.Context)

	GetExamById(c *gin.Context)
	UpdateExam(c *gin.Context)
	UpdateExamStatus(c *gin.Context)
	DeleteExam(c *gin.Context)

	GetAllAdmins(c *gin.Context)

	MustExamMiddleware() gin.HandlerFunc

	CreateAdmin(c *gin.Context)
	GetAdminsByExamId(c *gin.Context)

	MustAdminMiddleware() gin.HandlerFunc

	DeleteAdmin(c *gin.Context)
}

var _ SuperDelivery = &superHandler{}

func NewSuperDelivery(usecase superUsecase.SuperUsecase) SuperDelivery {
	return &superHandler{
		Usecase: usecase,
	}
}

type superHandler struct {
	Usecase superUsecase.SuperUsecase
}
