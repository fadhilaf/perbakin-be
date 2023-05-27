package route

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, delivery adminDelivery.AdminDelivery) {
	router.POST("/login", delivery.AdminLogin)

	adminRouter := router.Group("/", delivery.MustAdminMiddleware())
	adminRouter.GET("/", delivery.CheckAdminLogin)

	adminRouter.GET("/scorers", delivery.GetAllScorers)
	adminRouter.GET("/shooters", delivery.GetAllShooters)

	adminRouter.GET("/scorer", delivery.GetScorersByExamId)
	adminRouter.POST("/scorer", delivery.CreateScorer)

	adminRouter.GET("/shooter", delivery.GetShootersByExamId)

	scorerRouter := adminRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

	scorerRouter.GET("/", delivery.GetScorerById)
	scorerRouter.PUT("/", delivery.UpdateScorer)
	scorerRouter.DELETE("/", delivery.DeleteScorer)

	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId)
	scorerRouter.POST("/shooter", delivery.CreateShooter)

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", delivery.MustShooterMiddleware())

	shooterRouter.GET("/", delivery.GetShooterById)
	shooterRouter.PUT("/", delivery.UpdateShooter)
	shooterRouter.PUT("/image", delivery.UpdateShooterImage)
	shooterRouter.DELETE("/", delivery.DeleteShooter)
}
