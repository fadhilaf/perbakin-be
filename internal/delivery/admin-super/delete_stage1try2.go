package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage1try2(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage1Relation)

	res := handler.Usecase.DeleteStage1try2(model.ByIdRequest{ID: stage1.ID})

	c.JSON(res.Status, res)
}
