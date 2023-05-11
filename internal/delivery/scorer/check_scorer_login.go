package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) CheckScorerLogin(c *gin.Context) {
	scorerRelation := c.MustGet("scorer").(model.OperatorRelation)

	res := handler.Usecase.GetScorerByUserId(model.UserByUserIdRequest{UserID: scorerRelation.UserID})

	c.JSON(res.Status, res)
}
