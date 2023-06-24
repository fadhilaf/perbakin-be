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

	stage1Router := Stage1AllRoutes(resultRouter, allDelivery)

	stage1Router.POST("/2", delivery.MustStage1ModifyMiddleware(), allDelivery.CreateStage1try2)

	stage1ModifyRouter := stage1Router.Group("/:try", delivery.MustStage1ModifyMiddleware())
	stage1ModifyRouter.PUT("/no/:no", delivery.UpdateStage1No)
	stage1ModifyRouter.PUT("/checkmarks", delivery.UpdateStage1Checkmarks)
	stage1ModifyRouter.PATCH("/next", delivery.UpdateStage1NextNo)
	stage1ModifyRouter.PATCH("/finish", delivery.UpdateStage1Finish)

	Stage2AllRoutes(resultRouter, allDelivery)

	// stage2ModifyRouter := stage2Router.Group("/:try", delivery.MustStage2ModifyMiddleware())
	// stage2ModifyRouter.PUT("/no/:no", delivery.UpdateStage2No)
	// stage2ModifyRouter.PUT("/checkmarks", delivery.UpdateStage2Checkmarks)
	// stage2ModifyRouter.PATCH("/next", delivery.UpdateStage2NextNo)
	// stage2ModifyRouter.PATCH("/finish", delivery.UpdateStage2Finish)

	// stage3ModifyRouter := stage3Router.Group("/:try", delivery.MustStage3ModifyMiddleware())
	// stage3ModifyRouter.PUT("/no/:no", delivery.UpdateStage3No)
	// stage3ModifyRouter.PUT("/checkmarks", delivery.UpdateStage3Checkmarks)
	// stage3ModifyRouter.PATCH("/next", delivery.UpdateStage3NextNo)
	// stage3ModifyRouter.PATCH("/finish", delivery.UpdateStage3Finish)

	// stage4ModifyRouter := stage4Router.Group("/:try", delivery.MustStage4ModifyMiddleware())
	// stage4ModifyRouter.PUT("/no/:no", delivery.UpdateStage4No)
	// stage4ModifyRouter.PUT("/checkmarks", delivery.UpdateStage4Checkmarks)
	// stage4ModifyRouter.PATCH("/next", delivery.UpdateStage4NextNo)
	// stage4ModifyRouter.PATCH("/finish", delivery.UpdateStage4Finish)

	// stage5ModifyRouter := stage5Router.Group("/:try", delivery.MustStage5ModifyMiddleware())
	// stage5ModifyRouter.PUT("/no/:no", delivery.UpdateStage5No)
	// stage5ModifyRouter.PUT("/checkmarks", delivery.UpdateStage5Checkmarks)
	// stage5ModifyRouter.PATCH("/next", delivery.UpdateStage5NextNo)
	// stage5ModifyRouter.PATCH("/finish", delivery.UpdateStage5Finish)

	// stage6ModifyRouter := stage6Router.Group("/:try", delivery.MustStage6ModifyMiddleware())
	// stage6ModifyRouter.PUT("/no/:no", delivery.UpdateStage6No)
	// stage6ModifyRouter.PUT("/checkmarks", delivery.UpdateStage6Checkmarks)
	// stage6ModifyRouter.PATCH("/next", delivery.UpdateStage6NextNo)
	// stage6ModifyRouter.PATCH("/finish", delivery.UpdateStage6Finish)
}
