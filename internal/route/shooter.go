package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for scorerRouter
func ShooterAllRoutes(scorerRouter *gin.RouterGroup, delivery allDelivery.AllDelivery) *gin.RouterGroup {
	scorerRouter.GET("/shooter", delivery.GetShootersByScorerId) //done

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", delivery.MustShooterMiddleware())
	shooterRouter.GET("/", delivery.GetShooterById) //done

	return shooterRouter
}

// for scorerRouter, return shooterRouter
func ShooterAdminSuperRoutes(scorerRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) {
	scorerRouter.POST("/shooter", delivery.CreateShooter) //done

	shooterRouter := scorerRouter.Group("/shooter/:shooter_id", allDelivery.MustShooterMiddleware())
	shooterRouter.PUT("/", delivery.UpdateShooter)           //done
	shooterRouter.PUT("/image", delivery.UpdateShooterImage) //done
	shooterRouter.DELETE("/", delivery.DeleteShooter)        //done
}
