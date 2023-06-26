package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteStage3try2(c *gin.Context) {
	stage3 := c.MustGet("stage3").(model.Stage123456RelationAndStatus)

	res := handler.Usecase.DeleteStage3try2(model.ByIdRequest{ID: stage3.ID})

	c.JSON(res.Status, res)
}
