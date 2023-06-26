package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for resultRouter return stage4Router
func Stage4AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	//might delete create stage later
	// resultRouter.POST("/stage4", delivery.CreateStage4)

	stage4Router := resultRouter.Group("/stage4", delivery.MustStage4Middleware())
	stage4Router.GET("", delivery.GetStage4ById)

	return stage4Router
}

// for stage4Router
func Stage4AdminSuperRoutes(stage4Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	stage4Router.POST("/2", allDelivery.CreateStage4try2)
	stage4Router.PUT("", delivery.UpdateStage4)
	stage4Router.PUT("/sign", delivery.UpdateStage4Signs)
	stage4Router.DELETE("", delivery.DeleteStage4)
	stage4Router.DELETE("/2", delivery.DeleteStage4try2)
	stage4Router.PATCH("", delivery.FinishStage4)
}
