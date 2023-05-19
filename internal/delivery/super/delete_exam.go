package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) DeleteExam(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	res := handler.Usecase.DeleteExam(model.ByIdRequest{ID: exam.ID})

	c.JSON(res.Status, res)
}
