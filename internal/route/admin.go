package route

import (
	adminDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, delivery adminDelivery.AdminDelivery, allDelivery allDelivery.AllDelivery, adminSuperDelivery adminSuperDelivery.AdminSuperDelivery) {
	router.POST("/login", delivery.AdminLogin)

	adminRouter := router.Group("/", delivery.MustAdminMiddleware())
	adminRouter.GET("/", delivery.CheckAdminLogin)

	adminRouter.GET("/scorers", delivery.GetAllScorers)
	adminRouter.GET("/shooters", delivery.GetAllShooters)

	ExamAdminSuperRoutes(adminRouter, adminSuperDelivery)

	AdminSuperRoutes(adminRouter, adminSuperDelivery, allDelivery)
}
