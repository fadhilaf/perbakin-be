package route

import (
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"
	scorerDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"

	"github.com/gin-gonic/gin"
)

func ScorerRoutes(router *gin.RouterGroup, delivery scorerDelivery.ScorerDelivery, allDelivery allDelivery.AllDelivery) {
	router.POST("/login", delivery.ScorerLogin)

	scorerRouter := router.Group("", delivery.MustScorerMiddleware())
	scorerRouter.GET("", delivery.CheckScorerLogin)

	shooterRouter := ShooterAllRoutes(scorerRouter, allDelivery)

	resultRouter := ResultAllRoutes(shooterRouter, allDelivery)

	stage0Router := Stage0AllRoutes(resultRouter, allDelivery)

	stage0ModifyRouter := stage0Router.Group("", delivery.MustStage0ModifyMiddleware())
	stage0ModifyRouter.PUT("/no/:no", delivery.UpdateStage0Series)
	stage0ModifyRouter.PUT("/checkmarks", delivery.UpdateStage0Checkmarks)
	stage0ModifyRouter.PATCH("/next", delivery.UpdateStage0NextSeries)
	stage0ModifyRouter.PATCH("/finish", delivery.UpdateStage0Finish)

	//stage 1
	stage1Router := Stage1AllRoutes(resultRouter, allDelivery)

	stage1Router.POST("/2", delivery.MustStage1try1FinishedMiddleware(), allDelivery.CreateStage1try2)

	stage1ModifyRouter := stage1Router.Group("/:try", delivery.MustStage1ModifyMiddleware())
	stage1ModifyRouter.PUT("/no/:no", delivery.UpdateStage1No)
	stage1ModifyRouter.PUT("/checkmarks", delivery.UpdateStage1Checkmarks)
	stage1ModifyRouter.PATCH("/next", delivery.UpdateStage1NextNo)
	stage1ModifyRouter.PATCH("/finish", delivery.UpdateStage1Finish)

	//stage 2
	stage2Router := Stage2AllRoutes(resultRouter, allDelivery)

	stage2Router.POST("/2", delivery.MustStage2try1FinishedMiddleware(), allDelivery.CreateStage2try2)

	stage2ModifyRouter := stage2Router.Group("/:try", delivery.MustStage2ModifyMiddleware())
	stage2ModifyRouter.PUT("/no/:no", delivery.UpdateStage2No)
	stage2ModifyRouter.PUT("/checkmarks", delivery.UpdateStage2Checkmarks)
	stage2ModifyRouter.PATCH("/next", delivery.UpdateStage2NextNo)
	stage2ModifyRouter.PATCH("/finish", delivery.UpdateStage2Finish)

	//stage 3
	stage3Router := Stage3AllRoutes(resultRouter, allDelivery)

	stage3Router.POST("/2", delivery.MustStage3try1FinishedMiddleware(), allDelivery.CreateStage3try2)

	stage3ModifyRouter := stage3Router.Group("/:try", delivery.MustStage3ModifyMiddleware())
	stage3ModifyRouter.PUT("/no/:no", delivery.UpdateStage3No)
	stage3ModifyRouter.PUT("/checkmarks", delivery.UpdateStage3Checkmarks)
	stage3ModifyRouter.PATCH("/next", delivery.UpdateStage3NextNo)
	stage3ModifyRouter.PATCH("/finish", delivery.UpdateStage3Finish)
}
