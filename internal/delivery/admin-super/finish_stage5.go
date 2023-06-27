package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage5(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)

	if result.Stage != string(model.Stage5TypeString) {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage 5, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage5(model.ByIdRequest{ID: stage5.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage 5 ke stage 6", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage 5 ke stage 6", http.StatusOK, nil))
}
