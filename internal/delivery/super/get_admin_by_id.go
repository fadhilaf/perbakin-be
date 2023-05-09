package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetAdminById(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	id, ok := util.GetIdParam(c, "admin_id")
	if !ok {
		return
	}

	res := handler.Usecase.GetAdminById(model.OperatorByIdRequest{
		ID:     id,
		ExamID: exam.ID,
	})

	c.JSON(res.Status, res)
}
