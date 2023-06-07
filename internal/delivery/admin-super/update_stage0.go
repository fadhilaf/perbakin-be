package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage0(c *gin.Context) {
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

	var req model.UpdateStage0BodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//check series
	if ok := util.CheckSeries(c, "stage0",
		req.Series1,
		req.Series2,
		req.Series3,
		req.Series4,
		req.Series5,
	); !ok {
		return
	}

	//check checkmarks
	if ok := util.CheckCheckmarks(c, req.Checkmarks, "stage0"); !ok {
		return
	}

	res := handler.Usecase.UpdateStage0(model.UpdateStage0Request{
		ID:         stage0.ID,
		Status:     req.Status,
		Series1:    util.IntArrayToScores(req.Series1),
		Series2:    util.IntArrayToScores(req.Series2),
		Series3:    util.IntArrayToScores(req.Series3),
		Series4:    util.IntArrayToScores(req.Series4),
		Series5:    util.IntArrayToScores(req.Series5),
		Checkmarks: util.BoolArrayToCheckmarks(req.Checkmarks),
	})

	c.JSON(res.Status, res)
}
