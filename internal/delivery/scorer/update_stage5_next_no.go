package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage5NextNo(c *gin.Context) {
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	res := handler.Usecase.UpdateStage5NextNo(model.ByIdAndTryRequest{ID: stage5.ID, Try: try})

	c.JSON(res.Status, res)
}
