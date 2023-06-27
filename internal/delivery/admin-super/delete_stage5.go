package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage5(c *gin.Context) {
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage5(model.ByIdRequest{ID: stage5.ID})

	c.JSON(res.Status, res)
}
