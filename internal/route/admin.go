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
	adminRouter.GET("/scorer/:scorer_id", delivery.GetScorerById)
	adminRouter.PUT("/scorer/:scorer_id", delivery.UpdateScorer)
	adminRouter.DELETE("/scorer/:scorer_id", delivery.DeleteScorer)

	adminRouter.GET("/shooter", delivery.GetShootersByExamId)

	scorerRouter := adminRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId)
	scorerRouter.POST("/shooter", delivery.CreateShooter)
	scorerRouter.GET("/shooter/:shooter_id", delivery.GetShooterById)
	scorerRouter.PUT("/shooter/:shooter_id", delivery.UpdateShooter)
	scorerRouter.DELETE("/shooter/:shooter_id", delivery.DeleteShooter)
}
