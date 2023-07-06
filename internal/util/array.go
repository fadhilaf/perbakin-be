package util

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FadhilAF/perbakin-be/internal/model"
	"github.com/FadhilAF/perbakin-be/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/gin-gonic/gin"
)

func CheckNumbers(c *gin.Context, stageType model.StageList, numbers ...[]int) bool {
	for _, value := range numbers {
		if ok := CheckScores(c, value, stageType); !ok {
			return false
		}
	}

	return true
}

func CheckCheckmarks(c *gin.Context, checkmarks []bool, stageType model.StageList) bool {
	stage := make(map[model.StageList]int)
	stage[model.Stage1Type] = 6 //{max jumlah centang}}
	stage[model.Stage2Type] = 3
	stage[model.Stage3Type] = 6
	stage[model.Stage4Type] = 3
	stage[model.Stage5Type] = 2
	stage[model.Stage6Type] = 3

	// Count the number of true values
	numTrue := 0

	for _, value := range checkmarks {
		if value {
			numTrue++
		}
	}

	// The maximum number of true values is $max
	if numTrue > stage[stageType] {
		res := ToWebServiceResponse("Total jumlah centang tidak boleh lebih dari "+fmt.Sprint(stage[stageType]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}

func CheckScores(c *gin.Context, arr []int, stageType model.StageList) bool {
	//nyimpen tipe tipe stage
	stage := make(map[model.StageList]int)
	stage[model.Stage0Type] = 10 //{jumlah tembakan}}
	stage[model.Stage1Type] = 2
	stage[model.Stage2Type] = 2
	stage[model.Stage3Type] = 2
	stage[model.Stage4Type] = 2
	stage[model.Stage5Type] = 2
	stage[model.Stage6Type] = 2

	// Count the number of true values
	sum := 0
	for _, val := range arr {
		sum += val
	}

	// The maximum number of true values is $max
	if sum > stage[stageType] {
		res := ToWebServiceResponse("Total jumlah tembakan per nomor tidak boleh lebih dari "+fmt.Sprint(stage[stageType]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}

func CheckDurations(c *gin.Context, durations ...[]int) bool {
	for _, value := range durations {
		if ok := CheckDuration(c, value); !ok {
			return false
		}
	}

	return true
}

func CheckDuration(c *gin.Context, duration []int) bool {
	if duration[0] > 59 || duration[1] > 59 {
		res := ToWebServiceResponse("Durasi menit(duration[0]) dan detik(duration[1]) tidak boleh lebih dari 59", http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}

func MultiplyByIndexSum(list []int) int {
	sum := 0

	for i := 0; i < len(list); i++ {
		sum += list[i] * i
	}

	return sum
}

func GetSeriesTotalScoreFromDb(store repository.Store, c *gin.Context, ID pgtype.UUID, series int) (int, bool) {
	var seriesString string
	var repositoryErr error

	switch series {
	case 1:
		seriesString, repositoryErr = store.GetStage0Series1(context.Background(), ID)
	case 2:
		seriesString, repositoryErr = store.GetStage0Series2(context.Background(), ID)
	case 3:
		seriesString, repositoryErr = store.GetStage0Series3(context.Background(), ID)
	case 4:
		seriesString, repositoryErr = store.GetStage0Series4(context.Background(), ID)
	case 5:
		seriesString, repositoryErr = store.GetStage0Series5(context.Background(), ID)
	}

	if repositoryErr != nil {
		res := ToWebServiceResponse("Terjadi error ketika mengambil seri ke-"+strconv.Itoa(series)+" babak kualifikasi: "+repositoryErr.Error(), http.StatusInternalServerError, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return 0, false
	}

	return MultiplyByIndexSum(ScoresToIntArray(seriesString)), true
}

func CheckCheckmarksAmountAndSeriesMinimumScoreWithSeries(c *gin.Context, ID pgtype.UUID, checkmarks []bool, seriesScore [][]int, stageType model.StageList) bool {
	// Count the number of true values
	numTrue := 0

	for index, value := range checkmarks {
		if value {
			seriesTotal := MultiplyByIndexSum(seriesScore[index])

			// minimal jumlah skor seri untuk diberi centang
			if seriesTotal < model.CheckmarksScoreCheckParameter[stageType][1] {
				res := ToWebServiceResponse("Total nilai seri ke-"+fmt.Sprint(index+1)+" babak kualifikasi tidak boleh kurang dari "+fmt.Sprint(model.CheckmarksScoreCheckParameter[stageType][1])+" untuk dicentang", http.StatusBadRequest, nil)
				c.JSON(res.Status, res)
				c.Abort()
				return false
			}

			numTrue++
		}
	}

	// The maximum number of true values is $max
	if numTrue > model.CheckmarksScoreCheckParameter[stageType][0] {
		res := ToWebServiceResponse("Total jumlah centang tidak boleh lebih dari "+fmt.Sprint(model.CheckmarksScoreCheckParameter[stageType][0]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}
