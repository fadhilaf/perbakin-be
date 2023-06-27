package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage6(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage6(model.ByIdRequest{ID: stage6.ID})

	//edge case kalau delete stage yang paling terakhir dibuat
	if result.Stage == string(model.Stage6TypeString) || result.Stage == string(model.AllStageFinishedTypeString) {
		if err := handler.Usecase.UpdateResultStage(model.UpdateResultStageRequest{
			ID:    result.ID,
			Stage: string(model.Stage5TypeString),
		}); err {
			resMsg := util.ToWebServiceResponse("Terjadi kesalahan ketika memundurkan ke stage 5", http.StatusInternalServerError, nil)
			c.JSON(resMsg.Status, resMsg)
			c.Abort()
			return
		}
	}

	c.JSON(res.Status, res)
}
