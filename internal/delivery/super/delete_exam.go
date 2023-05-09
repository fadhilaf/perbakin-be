package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) DeleteExam(c *gin.Context) {
	super, ok := util.MustGetSuper(c)
	if !ok {
		return
	}

	id, ok := util.GetIdParam(c)
	if !ok {
		return
	}

	res := handler.Usecase.DeleteExam(model.DeleteExamRequest{ID: id, SuperID: super.ID})

	c.JSON(res.Status, res)
}
