package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) MustAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := util.GetUserIdFromContext(c)
		if !userId.Valid {
			res := util.ToWebServiceResponse("User belum login", 401, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		res := handler.Usecase.GetAdminByUserId(model.GetByUserIdRequest{UserID: userId})
		if res.Status != http.StatusOK {
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("admin", res.Data["admin"].(model.Admin))
		c.Next()
	}
}
