package delivery

import (
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) GetAllAdmins(c *gin.Context) {
	res := handler.Usecase.GetAllAdmins()

	c.JSON(res.Status, res.Data)
}
