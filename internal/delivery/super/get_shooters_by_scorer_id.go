package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetShootersByScorerId(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	res := handler.AllUsecase.GetShootersByScorerId(model.ByScorerIdRequest{ScorerID: scorer.ID})

	c.JSON(res.Status, res)
}
