package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSeries(c *gin.Context, series [][]int, stageType string) bool {
	for _, value := range series {
		if ok := CheckScores(c, value, stageType); !ok {
			return false
		}
	}

	return true
}

func CheckCheckmarks(c *gin.Context, arr []bool, stageType string) bool {
	stage := make(map[string][2]int)
	stage["stage0"] = [2]int{5, 3} // {{panjang array}, {jumlah centang}}

	// The input slice should have a length of $length
	if len(arr) != stage[stageType][0] {
		res := ToWebServiceResponse("Panjang array 'checkmarks' harus "+fmt.Sprint(stage[stageType][0]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	// Count the number of true values
	numTrue := 0

	for _, value := range arr {
		if value {
			numTrue++
		}
	}

	// The maximum number of true values is $max
	if numTrue > stage[stageType][1] {
		res := ToWebServiceResponse("Total jumlah centang tidak boleh lebih dari "+fmt.Sprint(stage[stageType][1]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}

func CheckScores(c *gin.Context, arr []int, stageType string) bool {
	//nyimpen tipe tipe stage
	stage := make(map[string][2]int)
	stage["stage0"] = [2]int{11, 10} // {{panjang array}, {jumlah tembakan}}

	// The input slice should have a length of $length
	if len(arr) != stage[stageType][0] {
		res := ToWebServiceResponse("Panjang array seri harus "+fmt.Sprint(stage[stageType][0]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	// Count the number of true values
	sum := 0
	for _, val := range arr {
		sum += val
	}

	// The maximum number of true values is $max
	if sum > stage[stageType][1] {
		res := ToWebServiceResponse("Total jumlah tembakan per seri tidak boleh lebih dari "+fmt.Sprint(stage[stageType][1]), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}
