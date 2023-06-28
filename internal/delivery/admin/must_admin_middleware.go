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
			util.DeleteAuthStatusCookie(c)

			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		adminExamAndStatus, err := handler.Usecase.GetAdminExamRelationByUserId(model.UserByUserIdRequest{UserID: userId})
		if err != nil {
			res := util.ToWebServiceResponse("User bukan merupakan admin", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if !adminExamAndStatus.Active {
			res := util.ToWebServiceResponse("Ujian tidak aktif", http.StatusForbidden, nil)
			util.RemoveAuthSession(c)
			util.DeleteAuthStatusCookie(c)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("admin", model.OperatorRelation{ID: adminExamAndStatus.ID, UserID: adminExamAndStatus.UserID, ExamID: adminExamAndStatus.ExamID})
		c.Set("exam", model.ExamRelation{ID: adminExamAndStatus.ExamID, SuperID: adminExamAndStatus.SuperID})
		c.Next()
	}
}
