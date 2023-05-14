package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) GetScorerById(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	id, ok := util.GetIdParam(c, "scorer_id")
	if !ok {
		return
	}

	res := handler.AdminSuperUsecase.GetScorerById(model.OperatorByIdRequest{
		ID:     id,
		ExamID: admin.ExamID,
	})

	c.JSON(res.Status, res)
}
