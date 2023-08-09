package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) GetScorerById(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	res := handler.Usecase.GetScorerById(model.ByIdRequest{
		ID: scorer.ID,
	})

	c.JSON(res.Status, res)
}
