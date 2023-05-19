package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) CreateShooter(c *gin.Context) {
	scorer := c.MustGet("scorer").(model.OperatorRelation)

	var req model.CreateShooterBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.AdminSuperUsecase.CreateShooter(model.CreateShooterRequest{
		ScorerID: scorer.ID,
		Body:     req,
	})

	c.JSON(res.Status, res)
}
