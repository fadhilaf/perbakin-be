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
			res := util.ToWebServiceResponse("User belum login", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		admin, err := handler.SessionUsecase.GetAdminRelationByUserId(model.UserByUserIdRequest{UserID: userId})
		if err != nil {
			res := util.ToWebServiceResponse("User bukan merupakan admin", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("admin", admin)
		c.Next()
	}
}
