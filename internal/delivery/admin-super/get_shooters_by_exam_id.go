package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) GetShootersByExamId(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	res := handler.Usecase.GetShootersByExamId(model.ByExamIdRequest{
		ExamID: exam.ID,
	})

	c.JSON(res.Status, res.Data)
}
