package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage1(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage1(model.ByIdRequest{ID: stage1.ID})

	//edge case kalau delete stage yang paling terakhir dibuat
	if result.Stage == string(model.Stage1TypeString) {
		if err := handler.Usecase.UpdateResultStage(model.UpdateResultStageRequest{
			ID:    result.ID,
			Stage: string(model.Stage0TypeString),
		}); err {
			resMsg := util.ToWebServiceResponse("Terjadi kesalahan ketika memundurkan ke babak kualifikasi", http.StatusInternalServerError, nil)
			c.JSON(resMsg.Status, resMsg)
			c.Abort()
			return
		}
	}

	c.JSON(res.Status, res)
}
