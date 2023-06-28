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
			res := util.ToWebServiceResponse("User belum login", http.StatusUnauthorized, nil)
			util.DeleteAuthStatusCookie(c)

			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		scorerAndStatus, err := handler.Usecase.GetScorerRelationByUserId(model.UserByUserIdRequest{UserID: userId})
		if err != nil {
			res := util.ToWebServiceResponse("User bukan merupakan penguji", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if !scorerAndStatus.Active {
			res := util.ToWebServiceResponse("Ujian tidak aktif", http.StatusForbidden, nil)
			util.RemoveAuthSession(c)
			util.DeleteAuthStatusCookie(c)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("scorer", model.OperatorRelation{ID: scorerAndStatus.ID, ExamID: scorerAndStatus.ExamID, UserID: scorerAndStatus.UserID})
		c.Next()
	}
}
