package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustScorerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := util.GetUserIdFromContext(c)
		if !userId.Valid {
			res := util.ToWebServiceResponse("User belum login", 401, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		res := handler.Usecase.GetScorerByUserId(model.GetByUserIdRequest{UserID: userId})
		if res.Status != http.StatusOK {
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("scorer", res.Data["scorer"].(model.Scorer))
		c.Next()
	}
}
