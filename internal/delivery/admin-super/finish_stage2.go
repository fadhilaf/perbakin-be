package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage2(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)

	if result.Stage != string(model.Stage2TypeString) {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage 2, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage2(model.ByIdRequest{ID: stage2.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage 2 ke stage 3", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage 2 ke stage 3", http.StatusOK, nil))
}
