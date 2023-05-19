package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) DeleteAdmin(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	res := handler.Usecase.DeleteAdmin(model.UserByUserIdRequest{
		UserID: admin.UserID,
	})

	c.JSON(res.Status, res)
}
