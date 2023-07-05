package usecase

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (usecase *utilUsecaseImpl) CheckSeriesuMinimumCheckmarkScore(c *gin.Context, ID pgtype.UUID, checkmarks []bool, stageType model.StageList) bool {
	stage := make(map[model.StageList][2]int)
	stage[model.Stage1Type] = [2]int{5, 70} //{max jumlah centang, minimum total nilai per seri}

	// Count the number of true values
	numTrue := 0

	for index, value := range checkmarks {
		if value {
			seriesTotal, ok := getSeriesTotalScore(usecase, c, ID, index+1)
			if !ok {
				return false
			}

			// minimal jumlah skor seri untuk diberi centang
			if seriesTotal < stage[stageType][1] {
				res := util.ToWebServiceResponse("Total nilai seri ke-"+fmt.Sprint(index+1)+" babak kualifikasi tidak boleh kurang dari "+fmt.Sprint(stage[stageType][1])+" untuk dicentang", http.StatusBadRequest, nil)
				c.JSON(res.Status, res)
				c.Abort()
				return false
			}

			numTrue++
		}
	}

	// The maximum number of true values is $max
	if numTrue > stage[stageType][0] {
		res := util.ToWebServiceResponse("Total jumlah centang tidak boleh lebih dari "+fmt.Sprint(stage[stageType]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}

func getSeriesTotalScore(usecase *utilUsecaseImpl, c *gin.Context, ID pgtype.UUID, series int) (int, bool) {
	var seriesString string
	var repositoryErr error

	switch series {
	case 1:
		seriesString, repositoryErr = usecase.GetStage0Series1(context.Background(), ID)
	case 2:
		seriesString, repositoryErr = usecase.GetStage0Series2(context.Background(), ID)
	case 3:
		seriesString, repositoryErr = usecase.GetStage0Series3(context.Background(), ID)
	case 4:
		seriesString, repositoryErr = usecase.GetStage0Series4(context.Background(), ID)
	case 5:
		seriesString, repositoryErr = usecase.GetStage0Series5(context.Background(), ID)
	}

	if repositoryErr != nil {
		res := util.ToWebServiceResponse("Terjadi error ketika mengambil seri ke-"+strconv.Itoa(series)+" babak kualifikasi: "+repositoryErr.Error(), http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return 0, false
	}

	return multiplyByIndexSum(util.ScoresToIntArray(seriesString)), true
}

func multiplyByIndexSum(list []int) int {
	sum := 0

	for i := 0; i < len(list); i++ {
		sum += list[i] * i
	}

	return sum
}
