package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) CreateStage4try2(c *gin.Context) {
	stage4 := c.MustGet("stage4").(model.Stage123456RelationAndStatus)
	res := handler.Usecase.CreateStage4try2(model.ByIdRequest{
		ID: stage4.ID,
	})

	c.JSON(res.Status, res)
}
