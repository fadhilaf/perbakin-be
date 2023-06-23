package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetStage2ById(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456Relation)
	res := handler.Usecase.GetStage2ById(model.ByIdRequest{
		ID: stage2.ID,
	})

	c.JSON(res.Status, res)
}
