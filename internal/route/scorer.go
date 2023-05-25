package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"

	"github.com/gin-gonic/gin"
)

func ScorerRoutes(router *gin.RouterGroup, delivery delivery.ScorerDelivery) {
	router.POST("/login", delivery.ScorerLogin)

	scorerRouter := router.Group("/", delivery.MustScorerMiddleware())
	scorerRouter.GET("/", delivery.CheckScorerLogin)

	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId)

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", delivery.MustShooterMiddleware())
	shooterRouter.GET("/", delivery.GetShooterById)

	shooterRouter.POST("/result", delivery.CreateResult)

	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware())
	resultRouter.GET("/", delivery.GetResultById)

	resultRouter.POST("/stage0", delivery.CreateStage0)

	stage0Router := resultRouter.Group("/stage0", delivery.MustStage0Middleware())
	stage0Router.GET("/", delivery.GetStage0ById)
	stage0Router.PUT("/series/:series", delivery.UpdateStage0Series)
}
