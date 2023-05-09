package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) MustExamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		examId, ok := util.GetIdParam(c, "exam_id")
		if !ok {
			return
		}

		exam, err := handler.AllUsecase.GetExamRelationById(model.ByIdRequest{ID: examId})
		if err != nil {
			res := util.ToWebServiceResponse("Ujian tidak ditemukan", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if super, exist := c.Get("super"); exist {
			super := super.(model.SuperRelation)

			if super.ID != exam.SuperID {
				res := util.ToWebServiceResponse("Tidak memiliki akses ke ujian ini", http.StatusUnauthorized, nil)
				c.JSON(res.Status, res)
				c.Abort()
				return
			}
		}

		if admin, exist := c.Get("admin"); exist {
			admin := admin.(model.OperatorRelation)

			if admin.ExamID != exam.ID {
				res := util.ToWebServiceResponse("Tidak memiliki akses ke ujian ini", http.StatusUnauthorized, nil)
				c.JSON(res.Status, res)
				c.Abort()
				return
			}
		}

		c.Set("exam", exam)
		c.Next()
	}
}
