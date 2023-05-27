package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) MustStage0Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelation)

		stage0, err := handler.AllUsecase.GetStage0RelationByResultId(model.ByResultIdRequest{
			ResultID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Hasil ujian kualifikasi belum ada", http.StatusNotFound, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("stage0", stage0)
		c.Next()
	}
}
