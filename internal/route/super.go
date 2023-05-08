package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/super"

	"github.com/gin-gonic/gin"
)

func SuperRoutes(router *gin.RouterGroup, delivery delivery.SuperDelivery) {
	router.POST("/login", delivery.SuperLogin)

	superRouter := router.Group("/", delivery.MustSuperMiddleware())
	superRouter.GET("/", delivery.CheckSuperLogin)
	superRouter.POST("/admin", delivery.CreateAdmin)
	superRouter.GET("/admin/:id", delivery.GetAdmin)
	superRouter.GET("/admins", delivery.GetAllAdmins)
	superRouter.PUT("/admin/:id", delivery.UpdateAdmin)
	superRouter.DELETE("/admin/:id", delivery.DeleteAdmin)
}
