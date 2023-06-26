package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage3NextNo(c *gin.Context) {
	stage3 := c.MustGet("stage3").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	res := handler.Usecase.UpdateStage3NextNo(model.ByIdAndTryRequest{ID: stage3.ID, Try: try})

	c.JSON(res.Status, res)
}
