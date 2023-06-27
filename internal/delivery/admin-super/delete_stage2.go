package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage2(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage2(model.ByIdRequest{ID: stage2.ID})

	//edge case kalau delete stage yang paling terakhir dibuat
	if result.Stage == string(model.Stage2TypeString) {
		if err := handler.Usecase.UpdateResultStage(model.UpdateResultStageRequest{
			ID:    result.ID,
			Stage: string(model.Stage1TypeString),
		}); err {
			resMsg := util.ToWebServiceResponse("Terjadi kesalahan ketika memundurkan ke stage 1", http.StatusInternalServerError, nil)
			c.JSON(resMsg.Status, resMsg)
			c.Abort()
			return
		}
	}

	c.JSON(res.Status, res)
}
