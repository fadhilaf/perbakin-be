package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) GetShootersByExamId(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	res := handler.AdminSuperUsecase.GetShootersByExamId(model.ByExamIdRequest{
		ExamID: admin.ExamID,
	})

	c.JSON(res.Status, res)
}
