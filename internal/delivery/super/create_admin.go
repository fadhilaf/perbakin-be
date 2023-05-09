package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CreateAdmin(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	var req model.OperatorBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.CreateAdmin(model.CreateOperatorRequest{
		ExamID: exam.ID,
		Body:   req,
	})

	c.JSON(res.Status, res)
}
