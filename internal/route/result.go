package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for shooterRouter return resultRouter
func ResultAllRoutes(shooterRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	//might delete create result later
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
