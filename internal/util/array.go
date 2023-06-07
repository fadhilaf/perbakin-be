package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckNumbers(c *gin.Context, stageType string, series ...[]int) bool {
	for _, value := range series {
		if ok := CheckScores(c, value, stageType); !ok {
			return false
		}
	}

	return true
}

func CheckCheckmarks(c *gin.Context, arr []bool, stageType string) bool {
	stage := make(map[string]int)
	stage["stage0"] = 3 //{jumlah centang}}
	stage["stage1"] = 6

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

func CheckScores(c *gin.Context, arr []int, stageType string) bool {
	//nyimpen tipe tipe stage
	stage := make(map[string]int)
	stage["stage0"] = 10 //{jumlah tembakan}}
	stage["stage1"] = 10

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

func CheckDuration(c *gin.Context, duration []int) bool {
	if duration[0] > 59 || duration[1] > 59 {
		res := ToWebServiceResponse("Durasi menit(duration[0]) dan detik(duration[1]) tidak boleh lebih dari 59", http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}
