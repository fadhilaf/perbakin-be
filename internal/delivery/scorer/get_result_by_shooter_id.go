package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) GetResultByShooterId(c *gin.Context) {
	shooter := c.MustGet("shooter").(model.ShooterRelation)

	res := handler.AllUsecase.GetResultByShooterId(model.ByShooterIdRequest{ShooterID: shooter.ID})

	c.JSON(res.Status, res)
}
