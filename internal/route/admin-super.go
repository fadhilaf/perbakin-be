package route

import (
	delivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"

	"github.com/gin-gonic/gin"
)

func AdminSuperRoutes(router *gin.RouterGroup, delivery delivery.AdminSuperDelivery) {
	adminSuperRoute := router.Group("/", delivery.MustAdminSuperMiddleware())

	adminSuperRoute.GET("/scorers", delivery.GetAllScorers)

	examRoute := adminSuperRoute.Group("/exam/:exam_id", delivery.MustExamMiddleware())

	examRoute.GET("/scorer", delivery.GetScorersByExamId)
	examRoute.POST("/scorer", delivery.CreateScorer)
	examRoute.GET("/scorers", delivery.GetScorersByExamId)
	examRoute.GET("/scorer/:scorer_id", delivery.GetScorerById)
	examRoute.PUT("/scorer/:scorer_id", delivery.UpdateScorer)
	examRoute.DELETE("/scorer/:scorer_id", delivery.DeleteScorer)
}
