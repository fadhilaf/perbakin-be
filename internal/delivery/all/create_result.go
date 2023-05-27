package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *allHandler) CreateResult(c *gin.Context) {
	shooter := c.MustGet("shooter").(model.ShooterRelation)

	res := handler.Usecase.CreateResult(model.ByShooterIdRequest{
		ShooterID: shooter.ID,
	})

	c.JSON(res.Status, res)
}
