package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) UpdateShooter(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)
	shooter := c.MustGet("shooter").(model.ShooterRelation)

	var req model.UpdateShooterBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.AdminSuperUsecase.UpdateShooter(model.UpdateShooterRequest{ID: shooter.ID, ExamID: exam.ID, Body: req})

	c.JSON(res.Status, res)
}
