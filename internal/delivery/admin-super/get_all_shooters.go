package delivery

import (
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) GetAllShooters(c *gin.Context) {
	res := handler.Usecase.GetAllShooters()

	c.JSON(res.Status, res)
}
