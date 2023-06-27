package delivery

import (
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) FinishStage6(c *gin.Context) {
	result := c.MustGet("result").(model.ResultRelationAndStatus)
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)

	if result.Stage != string(model.Stage6TypeString) {
		res := util.ToWebServiceResponse("Anda tidak dapat menyelesaikan stage 6, karena sedang tidak mengisi stage ini", http.StatusForbidden, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	err := handler.Usecase.FinishStage6(model.ByIdRequest{ID: stage6.ID})
	if err != nil {
		res := util.ToWebServiceResponse("Gagal mengupdate stage 6 ke selesai", http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.ToWebServiceResponse("Berhasil mengupdate stage 6 ke selesai", http.StatusOK, nil))
}
