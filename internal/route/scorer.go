package route

import (
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"
	scorerDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"

	"github.com/gin-gonic/gin"
)

func ScorerRoutes(router *gin.RouterGroup, delivery scorerDelivery.ScorerDelivery, allDelivery allDelivery.AllDelivery) {
	router.POST("/login", delivery.ScorerLogin)

	scorerRouter := router.Group("/", delivery.MustScorerMiddleware())
	scorerRouter.GET("/", delivery.CheckScorerLogin)

	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId)

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", delivery.MustShooterMiddleware())
	shooterRouter.GET("/", delivery.GetShooterById)

	ResultAllRoutes(shooterRouter, allDelivery)

	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware())
	stage0Router := resultRouter.Group("/stage0", delivery.MustStage0Middleware())
	stage0ModifyRouter := stage0Router.Group("/", delivery.MustStage0ModifyMiddleware())
	stage0ModifyRouter.PUT("/series/:series", delivery.UpdateStage0Series)
	stage0ModifyRouter.PUT("/checkmarks", delivery.UpdateStage0Checkmarks)
	stage0ModifyRouter.PATCH("/next", delivery.UpdateStage0NextSeries)
	stage0ModifyRouter.PATCH("/finish", delivery.UpdateStage0Finish)
}
