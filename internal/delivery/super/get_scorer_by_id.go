package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetScorerById(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	id, ok := util.GetIdParam(c, "scorer_id")
	if !ok {
		return
	}

	res := handler.AdminSuperUsecase.GetScorerById(model.OperatorByIdRequest{
		ID:     id,
		ExamID: exam.ID,
	})

	c.JSON(res.Status, res)
}
