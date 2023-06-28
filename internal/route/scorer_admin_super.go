package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"

	"github.com/gin-gonic/gin"
)

// for adminRouter or examRouter (for super), return scorerRouter
func ScorerAdminSuperRoutes(adminExamRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) *gin.RouterGroup {
	adminExamRouter.POST("/scorer", delivery.CreateScorer)

	scorerRouter := adminExamRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

	scorerRouter.GET("", delivery.GetScorerById)
	scorerRouter.PUT("", delivery.UpdateScorer)
	scorerRouter.PUT("/image", delivery.UpdateScorerImage)
	scorerRouter.DELETE("", delivery.DeleteScorer)

	return scorerRouter
}
