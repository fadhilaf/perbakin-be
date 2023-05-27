package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetStage0ById(c *gin.Context) {
	stage0 := c.MustGet("stage0").(model.Stage0Relation)
	res := handler.Usecase.GetStage0ById(model.ByIdRequest{
		ID: stage0.ID,
	})

	c.JSON(res.Status, res)
}
