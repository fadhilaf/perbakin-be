package delivery

import (
	"net/http"
	"strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) MustStage4Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := c.MustGet("result").(model.ResultRelationAndStatus)

		stageInt, _ := strconv.Atoi(result.Stage)
		if stageInt < int(model.Stage4Type) {
			res := util.ToWebServiceResponse("Anda tidak dapat mengakses stage 4, karena masih di stage sebelumnya", http.StatusForbidden, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		stage4, err := handler.Usecase.GetStage4RelationByResultId(model.ByResultIdRequest{
			ResultID: result.ID,
		})
		if err != nil {
			res := util.ToWebServiceResponse("Gagal mengambil hasil ujian stage 4: "+err.Error(), http.StatusInternalServerError, nil)
			c.JSON(res.Status, res)
			c.Abort()
			return
		}

		c.Set("stage4", stage4)
		c.Next()
	}
}
