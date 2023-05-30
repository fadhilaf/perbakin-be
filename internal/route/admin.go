package route

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, delivery adminDelivery.AdminDelivery, allDelivery allDelivery.AllDelivery, adminSuperDelivery adminSuperDelivery.AdminSuperDelivery) {
	router.POST("/login", delivery.AdminLogin)

	adminRouter := router.Group("", delivery.MustAdminMiddleware()) //jangan pake "/" agek jadi redirect https://stackoverflow.com/a/72164763
	adminRouter.GET("", delivery.CheckAdminLogin)

	adminRouter.GET("/scorers", adminSuperDelivery.GetAllScorers)
	adminRouter.GET("/shooters", adminSuperDelivery.GetAllShooters)

	AdminSuperRoutes(adminRouter, adminSuperDelivery, allDelivery)
}
