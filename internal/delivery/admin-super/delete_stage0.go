package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage0(c *gin.Context) {
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

	res := handler.Usecase.DeleteStage0(model.ByIdRequest{ID: stage0.ID})

	c.JSON(res.Status, res)
}
