package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage2Checkmarks(c *gin.Context) {
	stage2 := c.MustGet("stage2").(model.Stage123456RelationAndStatus)
	try := c.MustGet("try").(string)

	var req model.UpdateStage246CheckmarksBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate checkmarks
	if ok := util.CheckCheckmarks(c, req.Checkmarks, model.Stage2Type); !ok {
		return
	}

	res := handler.Usecase.UpdateStage2Checkmarks(model.UpdateStage123456CheckmarksRequest{ID: stage2.ID, Try: try, Checkmarks: util.BoolArrayToCheckmarks(req.Checkmarks)})

	c.JSON(res.Status, res)
}
