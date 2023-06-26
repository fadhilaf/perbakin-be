package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetStage4ById(c *gin.Context) {
	stage4 := c.MustGet("stage4").(model.Stage123456RelationAndStatus)
	res := handler.Usecase.GetStage4ById(model.ByIdRequest{
		ID: stage4.ID,
	})

	c.JSON(res.Status, res)
}
