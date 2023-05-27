package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for shooterRouter
func ResultAllRoutes(shooterRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) {
	shooterRouter.POST("/result", delivery.CreateResult)

	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware())
	resultRouter.GET("/", delivery.GetResultById)

	resultRouter.POST("/stage0", delivery.CreateStage0)

	stage0Router := resultRouter.Group("/stage0", delivery.MustStage0Middleware())
	stage0Router.GET("/", delivery.GetStage0ById)
}

// for shooterRouter
func ResultAdminSuperRoutes(shooterRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware())
	resultRouter.PUT("/", delivery.UpdateResult)
	resultRouter.DELETE("/", delivery.DeleteResult)

	stage0Router := resultRouter.Group("/stage0", delivery.MustStage0Middleware())
	// stage0Router.PUT("/", delivery.UpdateStage0)
	stage0Router.DELETE("/", delivery.DeleteStage0)
}
