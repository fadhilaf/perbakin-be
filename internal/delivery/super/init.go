package delivery

import (
	allUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/all"
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

func NewSuperDelivery(usecase superUsecase.SuperUsecase, allUsecase allUsecase.AllUsecase) SuperDelivery {
	return &superHandler{
		Usecase:    usecase,
		AllUsecase: allUsecase,
	}
}

type superHandler struct {
	Usecase    superUsecase.SuperUsecase
	AllUsecase allUsecase.AllUsecase
}
