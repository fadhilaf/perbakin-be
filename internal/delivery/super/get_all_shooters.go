package delivery

import (
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetAllShooters(c *gin.Context) {
	res := handler.AdminSuperUsecase.GetAllShooters()

	c.JSON(res.Status, res)
}
