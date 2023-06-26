package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage2(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage2(model.ByIdRequest{ID: stage2.ID})

	c.JSON(res.Status, res)
}
