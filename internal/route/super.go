package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"
	superDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"

	"github.com/gin-gonic/gin"
)

func SuperRoutes(router *gin.RouterGroup, delivery superDelivery.SuperDelivery, allDelivery allDelivery.AllDelivery, adminSuperDelivery adminSuperDelivery.AdminSuperDelivery) {
	router.POST("/login", delivery.SuperLogin)

	superRouter := router.Group("/", delivery.MustSuperMiddleware())
	superRouter.GET("/", delivery.CheckSuperLogin)

	superRouter.GET("/exams", delivery.GetAllExams)
	superRouter.GET("/admins", delivery.GetAllAdmins)
	superRouter.GET("/scorers", adminSuperDelivery.GetAllScorers)
	superRouter.GET("/shooters", adminSuperDelivery.GetAllShooters)

	superRouter.GET("/exam", delivery.GetExamsBySuperId)
	superRouter.POST("/exam", delivery.CreateExam)

	examRouter := superRouter.Group("/exam/:exam_id", delivery.MustExamMiddleware())
	examRouter.GET("/", delivery.GetExamById)
	examRouter.PUT("/", delivery.UpdateExam)
	examRouter.DELETE("/", delivery.DeleteExam)

	AdminSuperRoutes(examRouter, adminSuperDelivery, allDelivery)

	examRouter.GET("/admin", delivery.GetAdminsByExamId)
	examRouter.POST("/admin", delivery.CreateAdmin)

	adminRouter := examRouter.Group("/admin/:admin_id", delivery.MustAdminMiddleware())
	adminRouter.GET("/", delivery.GetAdminById)
	adminRouter.PUT("/", delivery.UpdateAdmin)
	adminRouter.DELETE("/", delivery.DeleteAdmin)

}
