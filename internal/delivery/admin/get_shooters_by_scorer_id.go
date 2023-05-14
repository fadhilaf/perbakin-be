package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) GetShootersByScorerId(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.Operator)

	res := handler.AllUsecase.GetShootersByScorerId(model.ByScorerIdRequest{ScorerID: scorer.ID})

	c.JSON(res.Status, res)
}
