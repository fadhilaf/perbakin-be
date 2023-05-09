package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) UpdateAdmin(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	id, ok := util.GetIdParam(c, "admin_id")
	if !ok {
		return
	}

	var req model.OperatorBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.UpdateAdmin(model.UpdateOperatorRequest{ID: id, ExamID: exam.ID, Body: req})

	c.JSON(res.Status, res)
}
