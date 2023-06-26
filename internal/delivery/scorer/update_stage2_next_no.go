package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage2NextNo(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	res := handler.Usecase.UpdateStage2NextNo(model.ByIdAndTryRequest{ID: stage2.ID, Try: try})

	c.JSON(res.Status, res)
}
