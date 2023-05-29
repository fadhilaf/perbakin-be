package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"

	"github.com/gin-gonic/gin"
)

// for adminRouter or examRouter (for super)
func ExamAdminSuperRoutes(adminExamRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	adminExamRouter.GET("/scorer", delivery.GetScorersByExamId)
	adminExamRouter.GET("/shooter", delivery.GetShootersByExamId)
}
