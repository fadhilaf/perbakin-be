package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage5(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage5(model.ByIdRequest{ID: stage5.ID})

	//edge case kalau delete stage yang paling terakhir dibuat
	if result.Stage == string(model.Stage5TypeString) {
		if err := handler.Usecase.UpdateResultStage(model.UpdateResultStageRequest{
			ID:    result.ID,
			Stage: string(model.Stage4TypeString),
		}); err {
			resMsg := util.ToWebServiceResponse("Terjadi kesalahan ketika memundurkan ke stage 4", http.StatusInternalServerError, nil)
			c.JSON(resMsg.Status, resMsg)
			c.Abort()
			return
		}
	}

	c.JSON(res.Status, res)
}
