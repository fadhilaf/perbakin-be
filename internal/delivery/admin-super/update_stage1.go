package delivery

import (
	"fmt"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
)

func (handler *adminSuperHandler) UpdateStage1(c *gin.Context) {
	// stage1 := c.MustGet("stage1").(model.Stage1Relation)

	var req model.UpdateStage13BodyRequest

	if ok := util.BindJSONAndValidate(c, &req); !ok {
		return
	}

	fmt.Println("status", req.Status, (req.Status.Status == ""))

	c.JSON(200, "yey")

	// //check series
	// if ok := util.CheckNumbers(c, "stage1",
	// 	req.Series1,
	// 	req.Series2,
	// 	req.Series3,
	// 	req.Series4,
	// 	req.Series5,
	// ); !ok {
	// 	return
	// }

	// //check checkmarks
	// if ok := util.CheckCheckmarks(c, req.Checkmarks, "stage1"); !ok {
	// 	return
	// }

	// res := handler.Usecase.UpdateStage1(model.UpdateStage1Request{
	// 	ID:         stage1.ID,
	// 	Status:     req.Status,
	// 	Series1:    util.IntArrayToScores(req.Series1),
	// 	Series2:    util.IntArrayToScores(req.Series2),
	// 	Series3:    util.IntArrayToScores(req.Series3),
	// 	Series4:    util.IntArrayToScores(req.Series4),
	// 	Series5:    util.IntArrayToScores(req.Series5),
	// 	Checkmarks: util.BoolArrayToCheckmarks(req.Checkmarks),
	// })

	// c.JSON(res.Status, res)
}
