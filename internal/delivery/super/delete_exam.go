package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) DeleteExam(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	res := handler.Usecase.DeleteExam(model.DeleteExamRequest{ID: exam.ID, SuperID: exam.SuperID})

	c.JSON(res.Status, res)
}
