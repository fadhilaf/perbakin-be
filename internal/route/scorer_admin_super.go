package route

import (
	adminSuperDelivery "github.com/FadhilAF/perbakin-be/internal/delivery/admin-super"

	"github.com/gin-gonic/gin"
)

// for adminRouter or examRouter (for super), return scorerRouter
func ScorerAdminSuperRoutes(adminExamRouter *gin.RouterGroup, delivery adminSuperDelivery.AdminSuperDelivery) *gin.RouterGroup {
	adminExamRouter.POST("/scorer", delivery.CreateScorer) //done

	scorerRouter := adminExamRouter.Group("/scorer/:scorer_id", delivery.MustScorerMiddleware())
	scorerRouter.GET("/", delivery.GetScorerById)   //done
	scorerRouter.PUT("/", delivery.UpdateScorer)    //done
	scorerRouter.DELETE("/", delivery.DeleteScorer) //done

	return scorerRouter
}
