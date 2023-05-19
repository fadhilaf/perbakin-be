package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, delivery delivery.AdminDelivery) {
	router.POST("/login", delivery.AdminLogin)

	adminRouter := router.Group("/", delivery.MustAdminMiddleware())
	adminRouter.GET("/", delivery.CheckAdminLogin)

	adminRouter.GET("/scorers", delivery.GetAllScorers)
	adminRouter.GET("/shooters", delivery.GetAllShooters)

	adminRouter.GET("/scorer", delivery.GetScorersByExamId)
	adminRouter.POST("/scorer", delivery.CreateScorer)

	adminRouter.GET("/shooter", delivery.GetShootersByExamId)

	scorerRouter := adminRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

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
