package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) DeleteShooter(c *gin.Context) {
	shooter := c.MustGet("shooter").(model.ShooterRelation)

	res := handler.AdminSuperUsecase.DeleteShooter(model.ByIdRequest{
		ID: shooter.ID,
	})

	c.JSON(res.Status, res)
}
