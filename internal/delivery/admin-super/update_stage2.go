package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage2(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)

	if stage2.IsTry2 {
		var req model.UpdateStage2try2BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try2
		if ok := util.CheckNumbers(c, model.Stage2Type,
			req.Try1.No1.Scores,
			req.Try1.No2.Scores,
			req.Try1.No3.Scores,
			req.Try2.No1.Scores,
			req.Try2.No2.Scores,
			req.Try2.No3.Scores,
		); !ok {
			return
		}

		//check duration try 2
		if ok := util.CheckDurations(c,
			req.Try2.No1.Duration,
			req.Try2.No2.Duration,
			req.Try2.No3.Duration,
			req.Try1.No1.Duration,
			req.Try1.No2.Duration,
			req.Try1.No3.Duration,
		); !ok {
			return
		}

		//check checkmarks try 1
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage2Type); !ok {
			return
		}
		//check checkmarks try 2
		if ok := util.CheckCheckmarksStage123456(c, req.Try2.Checkmarks, model.Stage2Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage2try2(model.UpdateStage246try2Request{
			ID: stage2.ID,
			Try1: model.Stage246TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.Scores, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.Scores, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try1.No3.Scores, req.Try1.No3.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
			Try2: model.Stage246TryString{
				Status:     req.Try2.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try2.No1.Scores, req.Try2.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try2.No2.Scores, req.Try2.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try2.No3.Scores, req.Try2.No3.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try2.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	} else {
		var req model.UpdateStage2try1BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try1
		if ok := util.CheckNumbers(c, model.Stage2Type,
			req.Try1.No1.Scores,
			req.Try1.No2.Scores,
			req.Try1.No3.Scores,
		); !ok {
			return
		}

		//check duration try 1
		if ok := util.CheckDurations(c,
			req.Try1.No1.Duration,
			req.Try1.No2.Duration,
			req.Try1.No3.Duration,
		); !ok {
			return
		}

		//check checkmarks
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage2Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage2try1(model.UpdateStage246try1Request{
			ID: stage2.ID,
			Try1: model.Stage246TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.Scores, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.Scores, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try1.No3.Scores, req.Try1.No3.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	}
}
