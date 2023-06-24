package delivery

import (
	"net/http"
	"strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) MustStage2Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelationAndStatus)

		stageInt, _ := strconv.Atoi(result.Stage)
		if stageInt < 2 {
			res := util.ToWebServiceResponse("Anda tidak dapat mengakses stage 2, karena masih di stage sebelumnya", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		stage2, err := handler.Usecase.GetStage2RelationByResultId(model.ByResultIdRequest{
			ResultID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Gagal mengambil hasil ujian stage 2: "+err.Error(), http.StatusInternalServerError, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("stage2", stage2)
		c.Next()
	}
}
