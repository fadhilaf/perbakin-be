package delivery

import (
	adminSuperUsecase "github.com/FadhilAF/perbakin-be/internal/usecase/admin-super"
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

	MustAdminMiddleware() gin.HandlerFunc

	GetAdminById(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	DeleteAdmin(c *gin.Context)
}

var _ SuperDelivery = &superHandler{}

func NewSuperDelivery(usecase superUsecase.SuperUsecase, adminSuperUsecase adminSuperUsecase.AdminSuperUsecase, allUsecase allUsecase.AllUsecase) SuperDelivery {
	return &superHandler{
		Usecase:           usecase,
		AdminSuperUsecase: adminSuperUsecase,
		AllUsecase:        allUsecase,
	}
}

type superHandler struct {
	Usecase           superUsecase.SuperUsecase
	AdminSuperUsecase adminSuperUsecase.AdminSuperUsecase
	AllUsecase        allUsecase.AllUsecase
}
