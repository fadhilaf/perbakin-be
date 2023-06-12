package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage1(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage1Relation)

	if stage1.IsTry2 {
		var req model.UpdateStage13try2BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try2
		if ok := util.CheckNumbers(c, "stage1",
			req.Try1.Scores1,
			req.Try1.Scores2,
			req.Try1.Scores3,
			req.Try1.Scores4,
			req.Try1.Scores5,
			req.Try1.Scores6,
			req.Try2.Scores1,
			req.Try2.Scores2,
			req.Try2.Scores3,
			req.Try2.Scores4,
			req.Try2.Scores5,
			req.Try2.Scores6,
		); !ok {
			return
		}

		//check duration try 2
		if ok := util.CheckDurations(c,
			req.Try2.Duration1,
			req.Try2.Duration2,
			req.Try2.Duration3,
			req.Try2.Duration4,
			req.Try2.Duration5,
			req.Try2.Duration6,
			req.Try1.Duration1,
			req.Try1.Duration2,
			req.Try1.Duration3,
			req.Try1.Duration4,
			req.Try1.Duration5,
			req.Try1.Duration6,
		); !ok {
			return
		}

		//check checkmarks try 1
		if ok := util.CheckCheckmarks(c, req.Try1.Checkmarks, "stage1"); !ok {
			return
		}
		//check checkmarks try 2
		if ok := util.CheckCheckmarks(c, req.Try2.Checkmarks, "stage1"); !ok {
			return
		}

		res := handler.Usecase.UpdateStage1try2(model.UpdateStage13try2Request{
			ID: stage1.ID,
			Try1: model.Stage13TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScores(req.Try1.Scores1, req.Try1.Duration1),
				No2:        util.IntArraysToScores(req.Try1.Scores2, req.Try1.Duration2),
				No3:        util.IntArraysToScores(req.Try1.Scores3, req.Try1.Duration3),
				No4:        util.IntArraysToScores(req.Try1.Scores4, req.Try1.Duration4),
				No5:        util.IntArraysToScores(req.Try1.Scores5, req.Try1.Duration5),
				No6:        util.IntArraysToScores(req.Try1.Scores6, req.Try1.Duration6),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
			Try2: model.Stage13TryString{
				Status:     req.Try2.Status,
				No1:        util.IntArraysToScores(req.Try2.Scores1, req.Try2.Duration1),
				No2:        util.IntArraysToScores(req.Try2.Scores2, req.Try2.Duration2),
				No3:        util.IntArraysToScores(req.Try2.Scores3, req.Try2.Duration3),
				No4:        util.IntArraysToScores(req.Try2.Scores4, req.Try2.Duration4),
				No5:        util.IntArraysToScores(req.Try2.Scores5, req.Try2.Duration5),
				No6:        util.IntArraysToScores(req.Try2.Scores6, req.Try2.Duration6),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try2.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	} else {
		var req model.UpdateStage13try1BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try1
		if ok := util.CheckNumbers(c, "stage1",
			req.Try1.Scores1,
			req.Try1.Scores2,
			req.Try1.Scores3,
			req.Try1.Scores4,
			req.Try1.Scores5,
			req.Try1.Scores6,
		); !ok {
			return
		}

		//check duration try 1
		if ok := util.CheckDurations(c,
			req.Try1.Duration1,
			req.Try1.Duration2,
			req.Try1.Duration3,
			req.Try1.Duration4,
			req.Try1.Duration5,
			req.Try1.Duration6,
		); !ok {
			return
		}

		//check checkmarks
		if ok := util.CheckCheckmarks(c, req.Try1.Checkmarks, "stage1"); !ok {
			return
		}

		res := handler.Usecase.UpdateStage1try1(model.UpdateStage13try1Request{
			ID: stage1.ID,
			Try1: model.Stage13TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScores(req.Try1.Scores1, req.Try1.Duration1),
				No2:        util.IntArraysToScores(req.Try1.Scores2, req.Try1.Duration2),
				No3:        util.IntArraysToScores(req.Try1.Scores3, req.Try1.Duration3),
				No4:        util.IntArraysToScores(req.Try1.Scores4, req.Try1.Duration4),
				No5:        util.IntArraysToScores(req.Try1.Scores5, req.Try1.Duration5),
				No6:        util.IntArraysToScores(req.Try1.Scores6, req.Try1.Duration6),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	}
}
