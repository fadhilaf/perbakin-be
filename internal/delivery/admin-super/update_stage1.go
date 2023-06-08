package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (handler *adminSuperHandler) UpdateStage1(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage1Relation)

	var req model.UpdateStage13BodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//check series
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

	//check duration
	if ok := util.CheckDurations(c,
		req.Try1.Duration1,
		req.Try1.Duration2,
		req.Try1.Duration3,
		req.Try1.Duration4,
		req.Try1.Duration5,
		req.Try1.Duration6,
		req.Try2.Duration1,
		req.Try2.Duration2,
		req.Try2.Duration3,
		req.Try2.Duration4,
		req.Try2.Duration5,
		req.Try2.Duration6,
	); !ok {
		return
	}

	//check checkmarks
	if ok := util.CheckCheckmarks(c, req.Try1.Checkmarks, "stage1"); !ok {
		return
	}
	if ok := util.CheckCheckmarks(c, req.Try2.Checkmarks, "stage1"); !ok {
		return
	}

	var (
		Try2Status     pgtype.Text
		Try2No1        pgtype.Text
		Try2No2        pgtype.Text
		Try2No3        pgtype.Text
		Try2No4        pgtype.Text
		Try2No5        pgtype.Text
		Try2No6        pgtype.Text
		Try2Checkmarks pgtype.Text
	)

	if stage1.IsTry2 {
		Try2Status.Scan(req.Try2.Status)
		Try2No1.Scan(util.IntArraysToScores(req.Try2.Scores1, req.Try2.Duration1))
		Try2No2.Scan(util.IntArraysToScores(req.Try2.Scores2, req.Try2.Duration2))
		Try2No3.Scan(util.IntArraysToScores(req.Try2.Scores3, req.Try2.Duration3))
		Try2No4.Scan(util.IntArraysToScores(req.Try2.Scores4, req.Try2.Duration4))
		Try2No5.Scan(util.IntArraysToScores(req.Try2.Scores5, req.Try2.Duration5))
		Try2No6.Scan(util.IntArraysToScores(req.Try2.Scores6, req.Try2.Duration6))
		Try2Checkmarks.Scan(util.BoolArrayToCheckmarks(req.Try2.Checkmarks))
	}

	res := handler.Usecase.UpdateStage1(model.UpdateStage13Request{
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
		Try2: model.Stage13TryStringOptional{
			Status:     Try2Status,
			No1:        Try2No1,
			No2:        Try2No2,
			No3:        Try2No3,
			No4:        Try2No4,
			No5:        Try2No5,
			No6:        Try2No6,
			Checkmarks: Try2Checkmarks,
		},
	})

	c.JSON(res.Status, res)
}
