package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage6(c *gin.Context) {
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)

	if stage6.IsTry2 {
		var req model.UpdateStage46try2BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try2
		if ok := util.CheckNumbers(c, model.Stage6Type,
			req.Try1.No1.ScoresA,
			req.Try1.No1.ScoresB,
			req.Try1.No2.ScoresA,
			req.Try1.No2.ScoresB,
			req.Try1.No3.ScoresA,
			req.Try1.No3.ScoresB,
			req.Try2.No1.ScoresA,
			req.Try2.No1.ScoresB,
			req.Try2.No2.ScoresA,
			req.Try2.No2.ScoresB,
			req.Try2.No3.ScoresA,
			req.Try2.No3.ScoresB,
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
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage6Type); !ok {
			return
		}
		//check checkmarks try 2
		if ok := util.CheckCheckmarksStage123456(c, req.Try2.Checkmarks, model.Stage6Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage6try2(model.UpdateStage246try2Request{
			ID: stage6.ID,
			Try1: model.Stage246TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.ScoresA, req.Try1.No1.ScoresB, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.ScoresA, req.Try1.No2.ScoresB, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try1.No3.ScoresA, req.Try1.No3.ScoresB, req.Try1.No3.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
			Try2: model.Stage246TryString{
				Status:     req.Try2.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try2.No1.ScoresA, req.Try2.No1.ScoresB, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try2.No2.ScoresA, req.Try2.No2.ScoresB, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try2.No3.ScoresA, req.Try2.No3.ScoresB, req.Try1.No3.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try2.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	} else {
		var req model.UpdateStage46try1BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try1
		if ok := util.CheckNumbers(c, model.Stage6Type,
			req.Try1.No1.ScoresA,
			req.Try1.No1.ScoresB,
			req.Try1.No2.ScoresA,
			req.Try1.No2.ScoresB,
			req.Try1.No3.ScoresA,
			req.Try1.No3.ScoresB,
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
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage6Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage6try1(model.UpdateStage246try1Request{
			ID: stage6.ID,
			Try1: model.Stage246TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.ScoresA, req.Try1.No1.ScoresB, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.ScoresA, req.Try1.No2.ScoresB, req.Try1.No2.Duration),
				No3:        util.IntArraysToScoresAndDuration(req.Try1.No3.ScoresA, req.Try1.No3.ScoresB, req.Try1.No3.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	}
}
