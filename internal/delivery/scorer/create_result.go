package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) CreateResult(c *gin.Context) {
	shooter := c.MustGet("shooter").(model.ShooterRelation)

	res := handler.AllUsecase.CreateResult(model.ByShooterIdRequest{
		ShooterID: shooter.ID,
	})

	c.JSON(res.Status, res)
}
