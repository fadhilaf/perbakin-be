package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage1Checkmarks(c *gin.Context) {
	stage1 := c.MustGet("stage1").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var req model.UpdateStage13CheckmarksBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate checkmarks
	if ok := util.CheckCheckmarksStage123456(c, req.Checkmarks, model.Stage1Type); !ok {
		return
	}

	res := handler.Usecase.UpdateStage1Checkmarks(model.UpdateStage123456CheckmarksRequest{ID: stage1.ID, Try: try, Checkmarks: util.BoolArrayToCheckmarks(req.Checkmarks)})

	c.JSON(res.Status, res)
}
