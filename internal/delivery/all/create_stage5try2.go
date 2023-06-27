package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) CreateStage5try2(c *gin.Context) {
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)
	res := handler.Usecase.CreateStage5try2(model.ByIdRequest{
		ID: stage5.ID,
	})

	c.JSON(res.Status, res)
}
