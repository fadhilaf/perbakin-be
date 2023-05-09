package delivery

import (
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetAllExams(c *gin.Context) {
	res := handler.Usecase.GetAllExams()

	c.JSON(res.Status, res.Data)
}
