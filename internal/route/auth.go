package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, delivery delivery.AuthDelivery) {
	router.POST("/logout", delivery.Logout)
}
