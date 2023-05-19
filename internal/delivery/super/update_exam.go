package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) UpdateExam(c *gin.Context) {
	exam := c.MustGet("exam").(model.ExamRelation)

	var req model.UpdateExamBodyStringRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.UpdateExam(model.UpdateExamRequest{ID: exam.ID, Body: model.UpdateExamBodyRequest{
		Name:      req.Name,
		Location:  req.Location,
		Organizer: req.Organizer,
		Begin:     util.ConvertDate(req.Begin),
		Finish:    util.ConvertDate(req.Finish),
	}})

	c.JSON(res.Status, res)
}
