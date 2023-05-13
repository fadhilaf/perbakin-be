package delivery

import (
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) GetAllScorers(c *gin.Context) {
	res := handler.Usecase.GetAllScorers()

	c.JSON(res.Status, res)
}
