package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, handler delivery.AdminDelivery) {
	router.POST("/login", handler.UserLogin)
}
