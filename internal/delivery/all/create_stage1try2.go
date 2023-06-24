package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) CreateStage1try2(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)
	res := handler.Usecase.CreateStage1try2(model.ByIdRequest{
		ID: stage1.ID,
	})

	c.JSON(res.Status, res)
}
