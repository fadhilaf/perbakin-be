package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage6(c *gin.Context) {
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage6(model.ByIdRequest{ID: stage6.ID})

	c.JSON(res.Status, res)
}
