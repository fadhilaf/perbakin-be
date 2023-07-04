package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage5Checkmarks(c *gin.Context) {
	stage5 := c.MustGet("stage5").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var req model.UpdateStage5CheckmarksBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate checkmarks
	if ok := util.CheckCheckmarksStage123456(c, req.Checkmarks, model.Stage5Type); !ok {
		return
	}

	res := handler.Usecase.UpdateStage5Checkmarks(model.UpdateStage123456CheckmarksRequest{ID: stage5.ID, Try: try, Checkmarks: util.BoolArrayToCheckmarks(req.Checkmarks)})

	c.JSON(res.Status, res)
}
