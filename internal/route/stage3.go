package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for resultRouter return stage3Router
func Stage3AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	//might delete create stage later
	// resultRouter.POST("/stage3", delivery.CreateStage3)

	stage3Router := resultRouter.Group("/stage3", delivery.MustStage3Middleware())
	stage3Router.GET("", delivery.GetStage3ById)

	return stage3Router
}

// for stage3Router
func Stage3AdminSuperRoutes(stage3Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	stage3Router.POST("/2", allDelivery.CreateStage3try2)
	stage3Router.PUT("", delivery.UpdateStage3)
	stage3Router.PUT("/sign", delivery.UpdateStage3Signs)
	stage3Router.DELETE("", delivery.DeleteStage3)
	stage3Router.DELETE("/2", delivery.DeleteStage3try2)
	stage3Router.PATCH("", delivery.FinishStage3)
}
