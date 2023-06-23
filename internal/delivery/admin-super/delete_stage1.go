package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage1(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage123456Relation)

	res := handler.Usecase.DeleteStage1(model.ByIdRequest{ID: stage1.ID})

	c.JSON(res.Status, res)
}
