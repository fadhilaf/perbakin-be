package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage6NextNo(c *gin.Context) {
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	res := handler.Usecase.UpdateStage6NextNo(model.ByIdAndTryRequest{ID: stage6.ID, Try: try})

	c.JSON(res.Status, res)
}
