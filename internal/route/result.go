package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for shooterRouter return resultRouter
func ResultAllRoutes(shooterRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	shooterRouter.POST("/result", delivery.CreateResult) //done

	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware()) //done
	resultRouter.GET("/", delivery.GetResultById)                                   //done

	return resultRouter
}

// for resultRouter
func ResultAdminSuperRoutes(resultRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	resultRouter.PUT("/", delivery.UpdateResult)    //done
	resultRouter.DELETE("/", delivery.DeleteResult) //done
}

// for resultRouter return stage0Router
func Stage0AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	resultRouter.POST("/stage0", delivery.CreateStage0)

	stage0Router := resultRouter.Group("/stage0", delivery.MustStage0Middleware())
	stage0Router.GET("/", delivery.GetStage0ById)

	return stage0Router
}

// for stage0Router
func Stage0AdminSuperRoutes(stage0Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	stage0Router.PUT("/", delivery.UpdateStage0)
	stage0Router.PUT("/sign", delivery.UpdateStage0Signs)
	stage0Router.DELETE("/", delivery.DeleteStage0)
}