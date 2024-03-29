package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustStage1ModifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelationAndStatus)
		stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)

		if result.Failed {
			res := util.ToWebServiceResponse("Hasil ujian tidak lulus", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if result.Stage != string(model.Stage1TypeString) {
			res := util.ToWebServiceResponse("Tidak dapat mengubah stage 1, sekarang sedang mengisi stage "+result.Stage, http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		var uri model.Stage123456TryRequestParam

		if ok := util.BindURIAndValidate(c, &uri); !ok {
			return
		}

		if !stage1.IsTry2 && uri.Try != "1" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah stage 1 percobaan 2, sekarang sedang mengisi percobaan 1", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if stage1.IsTry2 && uri.Try != "2" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah stage 1 percobaan 1, sekarang sedang mengisi percobaan 2", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("try", uri.Try)

		c.Next()
	}
}
