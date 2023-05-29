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

		adminExam, err := handler.Usecase.GetAdminExamRelationByUserId(model.UserByUserIdRequest{UserID: userId})
		if err != nil {
			res := util.ToWebServiceResponse("User bukan merupakan admin", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("admin", model.OperatorRelation{ID: adminExam.ID, UserID: adminExam.UserID, ExamID: adminExam.ExamID})
		c.Set("exam", model.ExamRelation{ID: adminExam.ExamID, SuperID: adminExam.SuperID})
		c.Next()
	}
}
