package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for scorerRouter
func ShooterAllRoutes(scorerRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId)

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", delivery.MustShooterMiddleware())
	shooterRouter.GET("/", delivery.GetShooterById)

	return shooterRouter
}

// for scorerRouter, return shooterRouter
func ShooterAdminSuperRoutes(scorerRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	scorerRouter.POST("/shooter", delivery.CreateShooter)

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", allDelivery.MustShooterMiddleware())
	shooterRouter.PUT("/", delivery.UpdateShooter)
	shooterRouter.PUT("/image", delivery.UpdateShooterImage)
	shooterRouter.DELETE("/", delivery.DeleteShooter)
}
