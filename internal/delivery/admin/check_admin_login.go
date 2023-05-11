package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) CheckAdminLogin(c *gin.Context) {
	adminRelation := c.MustGet("admin").(model.OperatorRelation)

	res := handler.Usecase.GetAdminByUserId(model.UserByUserIdRequest{UserID: adminRelation.UserID})

	c.JSON(res.Status, res)
}
