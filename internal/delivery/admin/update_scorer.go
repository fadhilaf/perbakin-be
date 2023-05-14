package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) UpdateScorer(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	id, ok := util.GetIdParam(c, "scorer_id")
	if !ok {
		return
	}

	var req model.OperatorBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.AdminSuperUsecase.UpdateScorer(model.UpdateOperatorRequest{ID: id, ExamID: admin.ExamID, Body: req})

	c.JSON(res.Status, res)
}
