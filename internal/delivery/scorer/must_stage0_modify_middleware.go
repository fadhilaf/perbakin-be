package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustStage0ModifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelationAndStatus)

		if result.Failed {
			res := util.ToWebServiceResponse("Hasil ujian tidak lulus", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if result.Stage != "0" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah babak kualifikasi, sekarang sedang mengisi babak "+result.Stage, http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
