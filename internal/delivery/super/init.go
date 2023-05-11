package delivery

import (
	sessionUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/session"
	superUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/super"

	"github.com/gin-gonic/gin"
)

type SuperDelivery interface {
	SuperLogin(c *gin.Context)

	MustSuperMiddleware() gin.HandlerFunc

	CheckSuperLogin(c *gin.Context)

	CreateExam(c *gin.Context)
	GetExamsBySuperId(c *gin.Context)
	GetAllExams(c *gin.Context)
	GetExamById(c *gin.Context)
	UpdateExam(c *gin.Context)
	DeleteExam(c *gin.Context)

	GetAllAdmins(c *gin.Context)

	MustExamMiddleware() gin.HandlerFunc

	CreateAdmin(c *gin.Context)
	GetAdminsByExamId(c *gin.Context)
	GetAdminById(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	DeleteAdmin(c *gin.Context)
}

var _ SuperDelivery = &superHandler{}

func NewSuperDelivery(usecase superUsecase.SuperUsecase, sessionUsecase sessionUsecase.SessionUsecase) SuperDelivery {
	return &superHandler{
		Usecase:        usecase,
		SessionUsecase: sessionUsecase,
	}
}

type superHandler struct {
	Usecase        superUsecase.SuperUsecase
	SessionUsecase sessionUsecase.SessionUsecase
}
