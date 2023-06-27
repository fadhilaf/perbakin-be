package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) CreateStage6try2(c *gin.Context) {
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)
	res := handler.Usecase.CreateStage6try2(model.ByIdRequest{
		ID: stage6.ID,
	})

	c.JSON(res.Status, res)
}
