package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustShooterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		scorer := c.MustGet("scorer").(model.OperatorRelation)

		shooterId, ok := util.GetIdParam(c, "shooter_id")
		if !ok {
			return
		}

		res := handler.AllUsecase.GetShooterById(shooterId)
		if res.Status != http.StatusOK {
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("shooter", res.Data["shooter"].(model.Shooter))
		c.Next()
	}
}
