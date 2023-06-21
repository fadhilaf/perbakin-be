package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for shooterRouter return resultRouter
func ResultAllRoutes(shooterRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	shooterRouter.POST("/result", delivery.CreateResult)

	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware())
	resultRouter.GET("", delivery.GetResultById)

	return resultRouter
}

// for resultRouter
func ResultAdminSuperRoutes(resultRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	resultRouter.PUT("", delivery.UpdateResult)
	resultRouter.DELETE("", delivery.DeleteResult)
}

// for resultRouter return stage0Router
func Stage0AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	resultRouter.POST("/stage0", delivery.CreateStage0)

	stage0Router := resultRouter.Group("/stage0", delivery.MustStage0Middleware())
	stage0Router.GET("", delivery.GetStage0ById)

	return stage0Router
}

// for stage0Router
func Stage0AdminSuperRoutes(stage0Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	stage0Router.PUT("", delivery.UpdateStage0)
	stage0Router.PUT("/sign", delivery.UpdateStage0Signs)
	stage0Router.DELETE("", delivery.DeleteStage0)
}

// for resultRouter return stage1Router
func Stage1AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	resultRouter.POST("/stage1", delivery.CreateStage1)

	stage1Router := resultRouter.Group("/stage1", delivery.MustStage1Middleware())
	stage1Router.GET("", delivery.GetStage1ById)

	return stage1Router
}

// for stage1Router
func Stage1AdminSuperRoutes(stage1Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	stage1Router.POST("/2", delivery.CreateStage1try2)
	stage1Router.PUT("", delivery.UpdateStage1)
	stage1Router.PUT("/sign", delivery.UpdateStage1Signs)
	stage1Router.DELETE("", delivery.DeleteStage1)
	stage1Router.DELETE("/2", delivery.DeleteStage1try2)
}
