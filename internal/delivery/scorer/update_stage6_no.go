package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage6No(c *gin.Context) {
	stage6 := c.MustGet("stage6").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var uri model.UpdateStage246NoUriRequest

	if ok := util.BindURIAndValidate(c, &uri); !ok {
		return
	}

	var req model.Stage46Numbers

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate scores
	if ok := util.CheckNumbers(c, model.Stage6Type, req.ScoresA, req.ScoresB); !ok {
		return
	}

	//validate duration
	if ok := util.CheckDuration(c, req.Duration); !ok {
		return
	}

	res := handler.Usecase.UpdateStage6No(model.UpdateStage123456NoRequest{
		ID:                stage6.ID,
		Try:               try,
		No:                uri.No,
		ScoresAndDuration: util.IntArraysToScoresAndDuration(req.ScoresA, req.ScoresB, req.Duration),
	})

	c.JSON(res.Status, res)
}
