package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for resultRouter return stage1Router
func Stage1AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	//might delete create stage later
	resultRouter.POST("/stage1", delivery.CreateStage1)

	stage1Router := resultRouter.Group("/stage1", delivery.MustStage1Middleware())
	stage1Router.GET("", delivery.GetStage1ById)

	return stage1Router
}

// for stage1Router
func Stage1AdminSuperRoutes(stage1Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	stage1Router.POST("/2", allDelivery.CreateStage1try2)
	stage1Router.PUT("", delivery.UpdateStage1)
	stage1Router.PUT("/sign", delivery.UpdateStage1Signs)
	stage1Router.DELETE("", delivery.DeleteStage1)
	stage1Router.DELETE("/2", delivery.DeleteStage1try2)
	stage1Router.PATCH("", delivery.FinishStage1)
}
