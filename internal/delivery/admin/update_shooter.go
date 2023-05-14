package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) UpdateShooter(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.Operator)

	id, ok := util.GetIdParam(c, "shooter_id")
	if !ok {
		return
	}

	var req model.UpdateShooterBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.AdminSuperUsecase.UpdateShooter(model.UpdateShooterRequest{ID: id, ExamID: scorer.ExamID, ScorerID: scorer.ID, Body: req})

	c.JSON(res.Status, res)
}
