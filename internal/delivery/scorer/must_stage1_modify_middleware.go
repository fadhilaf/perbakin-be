package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustStage1ModifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelation)
		stage1 := c.MustGet("stage1").(model.Stage123456Relation)

		status, err := handler.Usecase.GetResultStatusById(model.ByIdRequest{
			ID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Gagal mengambil hasil ujian: "+err.Error(), http.StatusInternalServerError, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if status.Failed {
			res := util.ToWebServiceResponse("Hasil ujian tidak lulus", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if status.Stage != "1" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah stage 1, sekarang sedang mengisi babak "+status.Stage, http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		var uri model.Stage123456TryRequestParam

		if ok := util.BindURIAndValidate(c, &uri); !ok {
			return
		}

		if !stage1.IsTry2 && uri.Try != "1" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah percobaan 2, sekarang sedang mengisi percobaan 1", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if stage1.IsTry2 && uri.Try != "2" {
			res := util.ToWebServiceResponse("Tidak dapat mengubah percobaan 1, sekarang sedang mengisi percobaan 2", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("try", uri.Try)

		c.Next()
	}
}
