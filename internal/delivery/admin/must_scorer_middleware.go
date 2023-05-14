package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) MustScorerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		admin := c.MustGet("admin").(model.OperatorRelation)

		scorerId, ok := util.GetIdParam(c, "scorer_id")
		if !ok {
			return
		}

		res := handler.AdminSuperUsecase.GetScorerById(model.OperatorByIdRequest{ID: scorerId, ExamID: admin.ExamID})
		if res.Status != http.StatusOK {
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("scorer", res.Data["scorer"].(model.Operator))
		c.Next()
	}
}
