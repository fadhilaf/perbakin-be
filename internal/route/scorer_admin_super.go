package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"
	allDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/all"

	"github.com/gin-gonic/gin"
)

// for adminRouter or examRouter (for super), return scorerRouter
func ScorerAdminSuperRoutes(adminExamRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery, allDelivery allDelivery.AllDelivery) *gin.RouterGroup {
	adminExamRouter.POST("/scorer", delivery.CreateScorer)

	scorerRouter := adminExamRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())

	scorerRouter.GET("", allDelivery.GetScorerById)
	scorerRouter.PUT("", allDelivery.UpdateScorer)
	scorerRouter.PUT("/image", allDelivery.UpdateScorerImage)
	scorerRouter.DELETE("", delivery.DeleteScorer)

	return scorerRouter
}
