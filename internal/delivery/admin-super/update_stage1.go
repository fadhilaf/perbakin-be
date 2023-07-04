package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage1(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)

	if stage1.IsTry2 {
		var req model.UpdateStage13try2BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try2
		if ok := util.CheckNumbers(c, model.Stage1Type,
			req.Try1.No1.Scores,
			req.Try1.No2.Scores,
			req.Try1.No3.Scores,
			req.Try1.No4.Scores,
			req.Try1.No5.Scores,
			req.Try1.No6.Scores,
			req.Try2.No1.Scores,
			req.Try2.No2.Scores,
			req.Try2.No3.Scores,
			req.Try2.No4.Scores,
			req.Try2.No5.Scores,
			req.Try2.No6.Scores,
		); !ok {
			return
		}

		//check duration try 2
		if ok := util.CheckDurations(c,
			req.Try2.No1.Duration,
			req.Try2.No2.Duration,
			req.Try2.No3.Duration,
			req.Try2.No4.Duration,
			req.Try2.No5.Duration,
			req.Try2.No6.Duration,
			req.Try1.No1.Duration,
			req.Try1.No2.Duration,
			req.Try1.No3.Duration,
			req.Try1.No4.Duration,
			req.Try1.No5.Duration,
			req.Try1.No6.Duration,
		); !ok {
			return
		}

		//check checkmarks try 1
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage1Type); !ok {
			return
		}
		//check checkmarks try 2
		if ok := util.CheckCheckmarksStage123456(c, req.Try2.Checkmarks, model.Stage1Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage1try2(model.UpdateStage13try2Request{
			ID: stage1.ID,
			Try1: model.Stage13TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.Scores, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.Scores, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try1.No3.Scores, req.Try1.No3.Duration),
				No4:        util.IntArraysToScoresAndDuration(req.Try1.No4.Scores, req.Try1.No4.Duration),
				No5:        util.IntArraysToScoresAndDuration(req.Try1.No5.Scores, req.Try1.No5.Duration),
				No6:        util.IntArraysToScoresAndDuration(req.Try1.No6.Scores, req.Try1.No6.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
			Try2: model.Stage13TryString{
				Status:     req.Try2.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try2.No1.Scores, req.Try2.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try2.No2.Scores, req.Try2.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try2.No3.Scores, req.Try2.No3.Duration),
				No4:        util.IntArraysToScoresAndDuration(req.Try2.No4.Scores, req.Try2.No4.Duration),
				No5:        util.IntArraysToScoresAndDuration(req.Try2.No5.Scores, req.Try2.No5.Duration),
				No6:        util.IntArraysToScoresAndDuration(req.Try2.No6.Scores, req.Try2.No6.Duration),
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
		if ok := util.CheckNumbers(c, model.Stage1Type,
			req.Try1.No1.Scores,
			req.Try1.No2.Scores,
			req.Try1.No3.Scores,
			req.Try1.No4.Scores,
			req.Try1.No5.Scores,
			req.Try1.No6.Scores,
		); !ok {
			return
		}

		//check duration try 1
		if ok := util.CheckDurations(c,
			req.Try1.No1.Duration,
			req.Try1.No2.Duration,
			req.Try1.No3.Duration,
			req.Try1.No4.Duration,
			req.Try1.No5.Duration,
			req.Try1.No6.Duration,
		); !ok {
			return
		}

		//check checkmarks
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage1Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage1try1(model.UpdateStage13try1Request{
			ID: stage1.ID,
			Try1: model.Stage13TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.Scores, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.Scores, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try1.No3.Scores, req.Try1.No3.Duration),
				No4:        util.IntArraysToScoresAndDuration(req.Try1.No4.Scores, req.Try1.No4.Duration),
				No5:        util.IntArraysToScoresAndDuration(req.Try1.No5.Scores, req.Try1.No5.Duration),
				No6:        util.IntArraysToScoresAndDuration(req.Try1.No6.Scores, req.Try1.No6.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	}
}
