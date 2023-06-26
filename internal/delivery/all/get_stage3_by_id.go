package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetStage3ById(c *gin.Context) {
	stage3 := c.MustGet("stage3").(model.Stage123456RelationAndStatus)
	res := handler.Usecase.GetStage3ById(model.ByIdRequest{
		ID: stage3.ID,
	})

	c.JSON(res.Status, res)
}
