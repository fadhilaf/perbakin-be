package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) MustExamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		super := c.MustGet("super").(model.SuperRelation)

		examId, ok := util.GetIdParam(c, "exam_id")
		if !ok {
			c.Abort()
			return
		}

		exam, err := handler.SessionUsecase.GetExamRelationById(model.ByIdRequest{ID: examId})
		if err != nil {
			res := util.ToWebServiceResponse("Ujian tidak ditemukan", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if super.ID != exam.SuperID {
			res := util.ToWebServiceResponse("Tidak memiliki akses ke ujian ini", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("exam", exam)
		c.Next()
	}
}
