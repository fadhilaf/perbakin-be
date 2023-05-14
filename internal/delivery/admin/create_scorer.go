package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminHandler) CreateScorer(c *gin.Context) {
	admin := c.MustGet("admin").(model.OperatorRelation)

	var req model.OperatorBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.AdminSuperUsecase.CreateScorer(model.CreateOperatorRequest{
		ExamID: admin.ExamID,
		Body:   req,
	})

	c.JSON(res.Status, res)
}
