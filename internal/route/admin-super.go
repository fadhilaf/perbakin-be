package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"

	"github.com/gin-gonic/gin"
)

func AdminSuperRoutes(router *gin.RouterGroup, delivery delivery.AdminSuperDelivery) {
	adminSuperRoute := router.Group("/", delivery.MustAdminSuperMiddleware())
	adminSuperRoute.POST("/scorer", delivery.CreateScorer)
	adminSuperRoute.GET("/scorers", delivery.GetAllScorers)
}
