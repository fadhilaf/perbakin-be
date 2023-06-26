package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage4(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage4 := c.MustGet("stage4").(model.Stage123456RelationAndStatus)

	if result.Stage != string(model.Stage4TypeString) {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage 4, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage4(model.ByIdRequest{ID: stage4.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage 4 ke stage 5", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage 4 ke stage 5", http.StatusOK, nil))
}
