package usecase

import (
	"fmt"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (usecase *utilUsecaseImpl) CheckCheckmarkAmountAndSeriesMinimumScore(c *gin.Context, ID pgtype.UUID, checkmarks []bool, stageType model.StageList) bool {
	// Count the number of true values
	numTrue := 0

	for index, value := range checkmarks {
		if value {
			seriesTotal, ok := util.GetSeriesTotalScoreFromDb(usecase.Store, c, ID, index+1)
			if !ok {
				return false
			}

			// minimal jumlah skor seri untuk diberi centang
			if seriesTotal < model.CheckmarksScoreCheckParameter[stageType][1] {
				res := util.ToWebServiceResponse("Total nilai seri ke-"+fmt.Sprint(index+1)+" babak kualifikasi tidak boleh kurang dari "+fmt.Sprint(model.CheckmarksScoreCheckParameter[stageType][1])+" untuk dicentang", http.StatusBadRequest, nil)
				c.JSON(res.Status, res)
				c.Abort()
				return false
			}

			numTrue++
		}
	}

	// The maximum number of true values is $max
	if numTrue > model.CheckmarksScoreCheckParameter[stageType][0] {
		res := util.ToWebServiceResponse("Total jumlah centang tidak boleh lebih dari "+fmt.Sprint(model.CheckmarksScoreCheckParameter[stageType][0]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}
