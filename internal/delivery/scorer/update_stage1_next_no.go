package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage1NextNo(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage1Relation)
	try := c.MustGet("try").(string)

	res := handler.Usecase.UpdateStage1NextNo(model.ByIdAndTryRequest{ID: stage1.ID, Try: try})

	c.JSON(res.Status, res)
}
