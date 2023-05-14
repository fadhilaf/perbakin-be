package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"

	"github.com/gin-gonic/gin"
)

func SuperRoutes(router *gin.RouterGroup, delivery delivery.SuperDelivery) {
	router.POST("/login", delivery.SuperLogin)

	superRouter := router.Group("/", delivery.MustSuperMiddleware())
	superRouter.GET("/", delivery.CheckSuperLogin)

	superRouter.GET("/exams", delivery.GetAllExams)
	superRouter.GET("/admins", delivery.GetAllAdmins)
	superRouter.GET("/scorers", delivery.GetAllScorers)
	superRouter.GET("/shooters", delivery.GetAllShooters)

	superRouter.GET("/exam", delivery.GetExamsBySuperId)
	superRouter.POST("/exam", delivery.CreateExam)

	examRouter := superRouter.Group("/exam/:exam_id", delivery.MustExamMiddleware())

	examRouter.GET("/", delivery.GetExamById)
	examRouter.PUT("/", delivery.UpdateExam)
	examRouter.DELETE("/", delivery.DeleteExam)

	examRouter.POST("/admin", delivery.CreateAdmin)
	examRouter.GET("/admin", delivery.GetAdminsByExamId)
	examRouter.GET("/admins", delivery.GetAllAdmins)
	examRouter.GET("/admin/:admin_id", delivery.GetAdminById)
	examRouter.PUT("/admin/:admin_id", delivery.UpdateAdmin)
	examRouter.DELETE("/admin/:admin_id", delivery.DeleteAdmin)

	examRouter.GET("/scorer", delivery.GetScorersByExamId)
	examRouter.POST("/scorer", delivery.CreateScorer)
	examRouter.GET("/scorer/:scorer_id", delivery.GetScorerById)
	examRouter.PUT("/scorer/:scorer_id", delivery.UpdateScorer)
	examRouter.DELETE("/scorer/:scorer_id", delivery.DeleteScorer)

	examRouter.GET("/shooter", delivery.GetShootersByExamId)

	scorerRoute := examRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

	scorerRoute.POST("/shooter", delivery.CreateShooter)

	scorerRoute.GET("/shooter/:shooter_id", delivery.GetShooterById)
	scorerRoute.PUT("/shooter/:shooter_id", delivery.UpdateShooter)
	scorerRoute.DELETE("/shooter/:shooter_id", delivery.DeleteShooter)
}
