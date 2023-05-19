package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) MustScorerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		exam := c.MustGet("exam").(model.ExamRelation)

		scorerId, ok := util.GetIdParam(c, "scorer_id")
		if !ok {
			c.Abort()
			return
		}

		scorer, err := handler.AdminSuperUsecase.GetScorerRelationById(model.ByIdRequest{ID: scorerId})
		if err != nil {
			res := util.ToWebServiceResponse("Penguji tidak ditemukan", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if scorer.ExamID != exam.ID {
			res := util.ToWebServiceResponse("Tidak dapat mengakses penguji ujian lain", http.StatusUnauthorized, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("scorer", scorer)
		c.Next()
	}
}
