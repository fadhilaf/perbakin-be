package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *superHandler) CreateExam(c *gin.Context) {
	super := c.MustGet("super").(model.SuperRelation)

	var req model.CreateExamBodyStringRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	res := handler.Usecase.CreateExam(model.CreateExamRequest{
		SuperID: super.ID,
		Body: model.CreateExamBodyRequest{
			Name:      req.Name,
			Location:  req.Location,
			Organizer: req.Organizer,
			Begin:     util.ConvertDate(req.Begin),
			Finish:    util.ConvertDate(req.Finish),
		},
	})

	c.JSON(res.Status, res)
}
