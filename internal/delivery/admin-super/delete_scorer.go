package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) DeleteScorer(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	res := handler.Usecase.DeleteScorer(model.UserByUserIdRequest{
		UserID: scorer.UserID,
	})

	c.JSON(res.Status, res)
}
