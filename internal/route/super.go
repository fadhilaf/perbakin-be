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

	examRouter.GET("/admin", delivery.GetAdminsByExamId)
	examRouter.POST("/admin", delivery.CreateAdmin)

	examRouter.GET("/scorer", delivery.GetScorersByExamId)
	examRouter.POST("/scorer", delivery.CreateScorer)

	examRouter.GET("/shooter", delivery.GetShootersByExamId)

	adminRouter := examRouter.Group("/admin/:admin_id", delivery.MustAdminMiddleware())

	adminRouter.GET("/admin/:admin_id", delivery.GetAdminById)
	adminRouter.PUT("/admin/:admin_id", delivery.UpdateAdmin)
	adminRouter.DELETE("/admin/:admin_id", delivery.DeleteAdmin)

	scorerRouter := examRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

	scorerRouter.GET("/scorer/:scorer_id", delivery.GetScorerById)
	scorerRouter.PUT("/scorer/:scorer_id", delivery.UpdateScorer)
	scorerRouter.DELETE("/scorer/:scorer_id", delivery.DeleteScorer)

	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId)
	scorerRouter.POST("/shooter", delivery.CreateShooter)

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", delivery.MustShooterMiddleware())

	shooterRouter.GET("/shooter/:shooter_id", delivery.GetShooterById)
	shooterRouter.PUT("/shooter/:shooter_id", delivery.UpdateShooter)
	shooterRouter.DELETE("/shooter/:shooter_id", delivery.DeleteShooter)
}
