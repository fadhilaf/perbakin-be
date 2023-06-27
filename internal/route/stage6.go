package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for resultRouter return stage6Router
func Stage6AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	//might delete create stage later
	// resultRouter.POST("/stage6", delivery.CreateStage6)

	stage6Router := resultRouter.Group("/stage6", delivery.MustStage6Middleware())
	stage6Router.GET("", delivery.GetStage6ById)

	return stage6Router
}

// for stage6Router
func Stage6AdminSuperRoutes(stage6Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	stage6Router.POST("/2", allDelivery.CreateStage6try2)
	stage6Router.PUT("", delivery.UpdateStage6)
	stage6Router.PUT("/sign", delivery.UpdateStage6Signs)
	stage6Router.DELETE("", delivery.DeleteStage6)
	stage6Router.DELETE("/2", delivery.DeleteStage6try2)
	stage6Router.PATCH("", delivery.FinishStage6)
}
