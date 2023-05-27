package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCheckmarks(c *gin.Context, arr []bool, length, max int) bool {
	// The input slice should have a length of $length
	if len(arr) != length {
		res := ToWebServiceResponse("Panjang array 'checkmarks' harus "+fmt.Sprint(length), http.StatusBadRequest, nil)
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
	if numTrue > 3 {
		res := ToWebServiceResponse("Total jumlah centang tidak boleh lebih dari "+fmt.Sprint(max), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}

func CheckScores(c *gin.Context, arr []int, length, max int) bool {
	// The input slice should have a length of $length
	if len(arr) != length {
		res := ToWebServiceResponse("Panjang array 'scores' harus "+fmt.Sprint(length), http.StatusBadRequest, nil)
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
	if sum > max {
		res := ToWebServiceResponse("Total jumlah tembakan per seri tidak boleh lebih dari "+fmt.Sprint(max), http.StatusBadRequest, nil)
		c.JSON(res.Status, res)
		c.Abort()
		return false
	}

	return true
}
