package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) GetAdminById(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	res := handler.Usecase.GetAdminById(model.ByIdRequest{
		ID: admin.ID,
	})

	c.JSON(res.Status, res)
}
