package delivery

import (
	"net/http"
	"strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage0(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

	stageInt, _ := strconv.Atoi(result.Stage)
	if stageInt != 0 {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage kualifikasi, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage0(model.ByIdRequest{ID: stage0.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage kualifikasi ke stage 1", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage kualifikasi ke stage 1", http.StatusOK, nil))
}
