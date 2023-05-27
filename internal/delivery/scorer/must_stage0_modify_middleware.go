package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustStage0ModifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelation)

		stage, err := handler.Usecase.GetResultStageById(model.ByIdRequest{
			ID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Hasil ujian belum ada", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if stage == "" {
			res := util.ToWebServiceResponse("Hasil ujian kualifikasi belum ada", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if stage != "0" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah babak kualifikasi, sekarang sedang mengisi babak "+stage, http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
