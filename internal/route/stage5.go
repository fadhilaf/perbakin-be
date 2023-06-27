package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for resultRouter return stage5Router
func Stage5AllRoutes(resultRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	//might delete create stage later
	// resultRouter.POST("/stage5", delivery.CreateStage5)

	stage5Router := resultRouter.Group("/stage5", delivery.MustStage5Middleware())
	stage5Router.GET("", delivery.GetStage5ById)

	return stage5Router
}

// for stage5Router
func Stage5AdminSuperRoutes(stage5Router *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	stage5Router.POST("/2", allDelivery.CreateStage5try2)
	stage5Router.PUT("", delivery.UpdateStage5)
	stage5Router.PUT("/sign", delivery.UpdateStage5Signs)
	stage5Router.DELETE("", delivery.DeleteStage5)
	stage5Router.DELETE("/2", delivery.DeleteStage5try2)
	stage5Router.PATCH("", delivery.FinishStage5)
}
