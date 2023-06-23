package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) MustStage2Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelation)

		stage2, err := handler.Usecase.GetStage2RelationByResultId(model.ByResultIdRequest{
			ResultID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Hasil ujian stage 1 belum ada", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("stage2", stage2)
		c.Next()
	}
}
