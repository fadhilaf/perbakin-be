package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) MustStage0Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelationAndStatus)

		stage0, err := handler.Usecase.GetStage0RelationByResultId(model.ByResultIdRequest{
			ResultID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Gagal mengambil hasil ujian kualifikasi: "+err.Error(), http.StatusInternalServerError, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("stage0", stage0)
		c.Next()
	}
}
