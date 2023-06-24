package delivery

import (
	"net/http"
	"strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage1(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)

	stageInt, _ := strconv.Atoi(result.Stage)
	if stageInt != 1 {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage 1, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage1(model.ByIdRequest{ID: stage1.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage 1 ke stage 2", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage 1 ke stage 2", http.StatusOK, nil))
}
