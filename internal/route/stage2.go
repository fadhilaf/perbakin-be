package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for resultRouter return stage2Router
func Stage2AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	resultRouter.POST("/stage2", delivery.CreateStage2)

	stage2Router := resultRouter.Group("/stage2", delivery.MustStage2Middleware())
	stage2Router.GET("", delivery.GetStage2ById)

	return stage2Router
}

// for stage2Router
func Stage2AdminSuperRoutes(stage2Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) {
	// stage2Router.POST("/2", delivery.CreateStage2try2)
	// stage2Router.PUT("", delivery.UpdateStage2)
	// stage2Router.PUT("/sign", delivery.UpdateStage2Signs)
	// stage2Router.DELETE("", delivery.DeleteStage2)
	// stage2Router.DELETE("/2", delivery.DeleteStage2try2)
}
