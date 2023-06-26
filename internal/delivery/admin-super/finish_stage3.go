package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage3(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage3 := c.MustGet("stage3").(model.Stage123456RelationAndStatus)

	if result.Stage != string(model.Stage3TypeString) {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage 3, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage3(model.ByIdRequest{ID: stage3.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage 3 ke stage 4", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage 3 ke stage 4", http.StatusOK, nil))
}
