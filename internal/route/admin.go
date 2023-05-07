package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, delivery delivery.AdminDelivery) {
	router.POST("/login", delivery.AdminLogin)

	superRouter := router.Group("/", delivery.MustAdminMiddleware())
	superRouter.GET("/", delivery.CheckAdminLogin)
}
