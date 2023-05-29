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

	shooterRouter := ShooterAllRoutes(scorerRouter, allDelivery)

	resultRouter := ResultAllRoutes(shooterRouter, allDelivery)

	stage0Router := Stage0AllRoutes(resultRouter, allDelivery)

	stage0ModifyRouter := stage0Router.Group("/", delivery.MustStage0ModifyMiddleware())
	stage0ModifyRouter.PUT("/series/:series", delivery.UpdateStage0Series)
	stage0ModifyRouter.PUT("/checkmarks", delivery.UpdateStage0Checkmarks)
	stage0ModifyRouter.PATCH("/next", delivery.UpdateStage0NextSeries)
	stage0ModifyRouter.PATCH("/finish", delivery.UpdateStage0Finish)

	// stage1Router := resultRouter.Group("/stage1", delivery.MustStage1Middleware())
	// stage1ModifyRouter := stage1Router.Group("/:try", delivery.MustStage1ModifyMiddleware())
	// stage1ModifyRouter.PUT("/series/:series", delivery.UpdateStage1Series)
	// stage1ModifyRouter.PUT("/checkmarks", delivery.UpdateStage1Checkmarks)
	// stage1ModifyRouter.PATCH("/next", delivery.UpdateStage1NextSeries)
	// stage1ModifyRouter.PATCH("/finish", delivery.UpdateStage1Finish)

	// stage2Router := resultRouter.Group("/stage2", delivery.MustStage2Middleware())
	// stage2ModifyRouter := stage2Router.Group("/:try", delivery.MustStage2ModifyMiddleware())
	// stage2ModifyRouter.PUT("/series/:series", delivery.UpdateStage2Series)
	// stage2ModifyRouter.PUT("/checkmarks", delivery.UpdateStage2Checkmarks)
	// stage2ModifyRouter.PATCH("/next", delivery.UpdateStage2NextSeries)
	// stage2ModifyRouter.PATCH("/finish", delivery.UpdateStage2Finish)

	// stage3Router := resultRouter.Group("/stage3", delivery.MustStage3Middleware())
	// stage3ModifyRouter := stage3Router.Group("/:try", delivery.MustStage3ModifyMiddleware())
	// stage3ModifyRouter.PUT("/series/:series", delivery.UpdateStage3Series)
	// stage3ModifyRouter.PUT("/checkmarks", delivery.UpdateStage3Checkmarks)
	// stage3ModifyRouter.PATCH("/next", delivery.UpdateStage3NextSeries)
	// stage3ModifyRouter.PATCH("/finish", delivery.UpdateStage3Finish)

	// stage4Router := resultRouter.Group("/stage4", delivery.MustStage4Middleware())
	// stage4ModifyRouter := stage4Router.Group("/:try", delivery.MustStage4ModifyMiddleware())
	// stage4ModifyRouter.PUT("/series/:series", delivery.UpdateStage4Series)
	// stage4ModifyRouter.PUT("/checkmarks", delivery.UpdateStage4Checkmarks)
	// stage4ModifyRouter.PATCH("/next", delivery.UpdateStage4NextSeries)
	// stage4ModifyRouter.PATCH("/finish", delivery.UpdateStage4Finish)

	// stage5Router := resultRouter.Group("/stage5", delivery.MustStage5Middleware())
	// stage5ModifyRouter := stage5Router.Group("/:try", delivery.MustStage5ModifyMiddleware())
	// stage5ModifyRouter.PUT("/series/:series", delivery.UpdateStage5Series)
	// stage5ModifyRouter.PUT("/checkmarks", delivery.UpdateStage5Checkmarks)
	// stage5ModifyRouter.PATCH("/next", delivery.UpdateStage5NextSeries)
	// stage5ModifyRouter.PATCH("/finish", delivery.UpdateStage5Finish)

	// stage6Router := resultRouter.Group("/stage6", delivery.MustStage6Middleware())
	// stage6ModifyRouter := stage6Router.Group("/:try", delivery.MustStage6ModifyMiddleware())
	// stage6ModifyRouter.PUT("/series/:series", delivery.UpdateStage6Series)
	// stage6ModifyRouter.PUT("/checkmarks", delivery.UpdateStage6Checkmarks)
	// stage6ModifyRouter.PATCH("/next", delivery.UpdateStage6NextSeries)
	// stage6ModifyRouter.PATCH("/finish", delivery.UpdateStage6Finish)
}
