package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage4NextNo(c *gin.Context) {
	stage4 := c.MustGet("stage4").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	res := handler.Usecase.UpdateStage4NextNo(model.ByIdAndTryRequest{ID: stage4.ID, Try: try})

	c.JSON(res.Status, res)
}
