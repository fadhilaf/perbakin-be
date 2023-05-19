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

	shooterRouter.GET("/result", delivery.GetResultByShooterId)
	shooterRouter.POST("/result", delivery.CreateResult)

	resultRouter := shooterRouter.Group("/result/:result_id", delivery.MustResultMiddleware())

	resultRouter.POST("/stage0", delivery.CreateStage0)
}
