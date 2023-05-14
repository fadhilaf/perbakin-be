package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) DeleteShooter(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.Operator)

	id, ok := util.GetIdParam(c, "shooter_id")
	if !ok {
		return
	}

	res := handler.AdminSuperUsecase.DeleteShooter(model.ShooterByIdRequest{
		ID:       id,
		ScorerID: scorer.ID,
	})

	c.JSON(res.Status, res)
}
