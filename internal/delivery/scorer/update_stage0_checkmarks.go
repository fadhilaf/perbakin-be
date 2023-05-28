package delivery

import (
	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *scorerHandler) UpdateStage0Checkmarks(c *gin.Context) {
	stage0 := c.MustGet("stage0").(model.Stage0Relation)

	var req model.UpdateStage0CheckmarksBodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	//validate checkmarks
	if ok := util.CheckCheckmarks(c, req.Checkmarks, "stage0"); !ok {
		return
	}

	res := handler.Usecase.UpdateStage0Checkmarks(model.UpdateStage0CheckmarksRequest{ID: stage0.ID, Checkmarks: util.BoolArrayToCheckmarks(req.Checkmarks)})

	c.JSON(res.Status, res)
}
