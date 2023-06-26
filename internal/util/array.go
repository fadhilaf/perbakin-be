package util

import (
	"fmt"
	"net/http"

	"github.com/FadhilAF/perbakin-be/internal/model"

	"github.com/gin-gonic/gin"
)

func CheckNumbers(c *gin.Context, stageType model.StageList, series ...[]int) bool {
	for _, value := range series {
		if ok := CheckScores(c, value, stageType); !ok {
			return false
		}
	}

	return true
}

func CheckCheckmarks(c *gin.Context, arr []bool, stageType model.StageList) bool {
	stage := make(map[model.StageList]int)
	stage[model.Stage0Type] = 3 //{jumlah centang}}
	stage[model.Stage1Type] = 6
	stage[model.Stage2Type] = 3
	stage[model.Stage3Type] = 6
	stage[model.Stage4Type] = 3
	stage[model.Stage5Type] = 2
	stage[model.Stage6Type] = 3

	// Count the number of true values
	numTrue := 0

	for _, value := range arr {
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
	stage[model.Stage1Type] = 10
	stage[model.Stage2Type] = 10
	stage[model.Stage3Type] = 10
	stage[model.Stage4Type] = 10
	stage[model.Stage5Type] = 10
	stage[model.Stage6Type] = 10

	// Count the number of true values
	sum := 0
	for _, val := range arr {
		sum += val
	}

	// The maximum number of true values is $max
	if sum > stage[stageType] {
		res := ToWebServiceResponse("Total jumlah tembakan per seri tidak boleh lebih dari "+fmt.Sprint(stage[stageType]), http.StatusBadRequest, nil)
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
