package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) MustStage1tryFinishedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)

		isTry1Finished, err := handler.Usecase.Stage1try1Finished(model.ByIdRequest{
			ID: stage1.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Gagal mengambil status stage 1: "+err.Error(), http.StatusInternalServerError, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		if !isTry1Finished {
			res := util.ToWebServiceResponse("Tidak dapat membuat stage 1 percobaan 2. Percobaan 1 belum selesai", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
