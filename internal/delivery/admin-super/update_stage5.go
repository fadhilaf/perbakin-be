package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage5(c *gin.Context) {
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)

	if stage5.IsTry2 {
		var req model.UpdateStage5try2BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try2
		if ok := util.CheckNumbers(c, model.Stage5Type,
			req.Try1.No1.ScoresA,
			req.Try1.No1.ScoresB,
			req.Try1.No1.ScoresC,
			req.Try1.No2.ScoresA,
			req.Try1.No2.ScoresB,
			req.Try1.No2.ScoresC,
			req.Try2.No1.ScoresA,
			req.Try2.No1.ScoresB,
			req.Try2.No1.ScoresC,
			req.Try2.No2.ScoresA,
			req.Try2.No2.ScoresB,
			req.Try2.No2.ScoresC,
		); !ok {
			return
		}

		//check duration try 2
		if ok := util.CheckDurations(c,
			req.Try2.No1.Duration,
			req.Try2.No2.Duration,
			req.Try1.No1.Duration,
			req.Try1.No2.Duration,
		); !ok {
			return
		}

		//check checkmarks try 1
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage5Type); !ok {
			return
		}
		//check checkmarks try 2
		if ok := util.CheckCheckmarksStage123456(c, req.Try2.Checkmarks, model.Stage5Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage5try2(model.UpdateStage5try2Request{
			ID: stage5.ID,
			Try1: model.Stage5TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.ScoresA, req.Try1.No1.ScoresB, req.Try1.No1.ScoresC, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.ScoresA, req.Try1.No2.ScoresB, req.Try1.No2.ScoresC, req.Try1.No2.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
			Try2: model.Stage5TryString{
				Status:     req.Try2.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try2.No1.ScoresA, req.Try2.No1.ScoresB, req.Try2.No1.ScoresC, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try2.No2.ScoresA, req.Try2.No2.ScoresB, req.Try2.No2.ScoresC, req.Try1.No2.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try2.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	} else {
		var req model.UpdateStage5try1BodyRequest

		if ok := util.BindJSONAndValidate(c, &req); !ok {
			return
		}

		//check series try1
		if ok := util.CheckNumbers(c, model.Stage5Type,
			req.Try1.No1.ScoresA,
			req.Try1.No1.ScoresB,
			req.Try1.No1.ScoresC,
			req.Try1.No2.ScoresA,
			req.Try1.No2.ScoresB,
			req.Try1.No2.ScoresC,
		); !ok {
			return
		}

		//check duration try 1
		if ok := util.CheckDurations(c,
			req.Try1.No1.Duration,
			req.Try1.No2.Duration,
		); !ok {
			return
		}

		//check checkmarks
		if ok := util.CheckCheckmarksStage123456(c, req.Try1.Checkmarks, model.Stage5Type); !ok {
			return
		}

		res := handler.Usecase.UpdateStage5try1(model.UpdateStage5try1Request{
			ID: stage5.ID,
			Try1: model.Stage5TryString{
				Status:     req.Try1.Status,
				No1:        util.IntArraysToScoresAndDuration(req.Try1.No1.ScoresA, req.Try1.No1.ScoresB, req.Try1.No1.ScoresC, req.Try1.No1.Duration),
				No2:        util.IntArraysToScoresAndDuration(req.Try1.No2.ScoresA, req.Try1.No2.ScoresB, req.Try1.No2.ScoresC, req.Try1.No2.Duration),
				Checkmarks: util.BoolArrayToCheckmarks(req.Try1.Checkmarks),
			},
		})

		c.JSON(res.Status, res)
	}
}
