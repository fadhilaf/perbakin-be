package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage5No(c *gin.Context) {
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var uri model.UpdateStage5NoUriRequest

	if ok := util.BindURIAndValidate(c, &uri); !ok {
		return
	}

	var req model.Stage5Numbers

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate scores
	if ok := util.CheckNumbers(c, model.Stage5Type, req.ScoresA, req.ScoresB, req.ScoresC); !ok {
		return
	}

	//validate duration
	if ok := util.CheckDuration(c, req.Duration); !ok {
		return
	}

	res := handler.Usecase.UpdateStage5No(model.UpdateStage123456NoRequest{
		ID:                stage5.ID,
		Try:               try,
		No:                uri.No,
		ScoresAndDuration: util.IntArraysToScoresAndDuration(req.ScoresA, req.ScoresB, req.ScoresC, req.Duration),
	})

	c.JSON(res.Status, res)
}
