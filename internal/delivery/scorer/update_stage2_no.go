package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage2No(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var uri model.UpdateStage246NoUriRequest

	if ok := util.BindURIAndValidate(c, &uri); !ok {
		return
	}

	var req model.Stage123Numbers

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate scores
	if ok := util.CheckScores(c, req.Scores, model.Stage2Type); !ok {
		return
	}

	//validate duration
	if ok := util.CheckDuration(c, req.Duration); !ok {
		return
	}

	res := handler.Usecase.UpdateStage2No(model.UpdateStage123456NoRequest{
		ID:                stage2.ID,
		Try:               try,
		No:                uri.No,
		ScoresAndDuration: util.IntArraysToScoresAndDuration(req.Scores, req.Duration),
	})

	c.JSON(res.Status, res)
}
