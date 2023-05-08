package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/scorer"

	"github.com/gin-gonic/gin"
)

func ScorerRoutes(router *gin.RouterGroup, delivery delivery.ScorerDelivery) {
	router.POST("/login", delivery.ScorerLogin)

	superRouter := router.Group("/", delivery.MustScorerMiddleware())
	superRouter.GET("/", delivery.CheckScorerLogin)
}
