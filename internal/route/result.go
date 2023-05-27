package route

import (
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for shooterRouter
func ResultAllRoutes(shooterRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) {
	shooterRouter.POST("/result", delivery.CreateResult)

	resultRouter := shooterRouter.Group("/result", delivery.MustResultMiddleware())
	resultRouter.GET("/", delivery.GetResultById)
	// resultRouter.POST("/stage0", delivery.CreateStage0)
}
