package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) GetExamById(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	res := handler.Usecase.GetExamById(model.ByIdRequest{ID: exam.ID})

	c.JSON(res.Status, res)
}
