package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetStage1ById(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage1Relation)
	res := handler.Usecase.GetStage1ById(model.ByIdRequest{
		ID: stage1.ID,
	})

	c.JSON(res.Status, res)
}
