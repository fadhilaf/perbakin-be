package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) MustAdminSuperMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := util.GetUserIdFromContext(c)
		if !userId.Valid {
			res := util.ToWebServiceResponse("User belum login", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if super, err := handler.SessionUsecase.GetSuperRelationByUserId(model.UserByUserIdRequest{UserID: userId}); err == nil {
			c.Set("super", super)
			c.Next()
			return
		}

		if admin, err := handler.SessionUsecase.GetAdminRelationByUserId(model.UserByUserIdRequest{UserID: userId}); err == nil {
			c.Set("admin", admin)
			c.Next()
			return
		}

		res := util.ToWebServiceResponse("User bukan admin maupun super admin", http.StatusUnauthorized, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}
}
