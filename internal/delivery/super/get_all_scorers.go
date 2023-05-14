package delivery

import (
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetAllScorers(c *gin.Context) {
	res := handler.AdminSuperUsecase.GetAllScorers()

	c.JSON(res.Status, res)
}
