package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage0Series(c *gin.Context) {
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

	var uri model.UpdateStage0SeriesUriRequest

	if ok := util.BindURIAndValidate(c, &uri); !ok {
		return
	}

	var req model.UpdateStage0SeriesBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate scores
	if ok := util.CheckScores(c, req.Scores, 11, 10); !ok {
		return
	}

	res := handler.Usecase.UpdateStage0Series(model.UpdateStage0SeriesRequest{ID: stage0.ID, Series: uri.Series, Scores: util.IntArrayToScores(req.Scores)})

	c.JSON(res.Status, res)
}
