package util

import (
	"strconv"
	"strings"

	"github.com/FadhilAF/perbakin-be/internal/model"
)

func Stage5DatabaseNumbersToStruct(input string) model.Stage5Numbers {
	array := NumbersToIntArrayArray(input)

	return model.Stage5Numbers{
		ScoresA:  array[0],
		ScoresB:  array[1],
		ScoresC:  array[2],
		Duration: array[3],
	}
}

func Stage46DatabaseNumbersToStruct(input string) model.Stage46Numbers {
	array := NumbersToIntArrayArray(input)

	return model.Stage46Numbers{
		ScoresA:  array[0],
		ScoresB:  array[1],
		Duration: array[2],
	}
}

func Stage123DatabaseNumbersToStruct(input string) model.Stage123Numbers {
	array := NumbersToIntArrayArray(input)

	return model.Stage123Numbers{
		Scores:   array[0],
		Duration: array[1],
	}
}

func NumbersToIntArrayArray(input string) [][]int {
	//menghilangkan simbol (, ), dan "
	inputTrimmed := strings.Trim(input, `()"`)

	//memisahkan string dari substring yang berbentuk ","
	nums := strings.Split(inputTrimmed, `","`)

	result := make([][]int, len(nums))

	for i, numStr := range nums {
		result[i] = ScoresToIntArray(numStr)
	}

	return result
}

// mengubah string `(0,0,0)` atau `0,0,0` atau `(0,0,0` atau `0,0,0)` jadi array [0,0,0]
func ScoresToIntArray(input string) []int {
	// ngilangkan kurung and dan memisahkan satu satu string yang dipisahkan oleh koma
	inputTrimmed := strings.Trim(input, "()")
	nums := strings.Split(inputTrimmed, ",")

	// Initialize the integer array
	result := make([]int, len(nums))

	// Convert each number from string to int and store in the array
	for i, numStr := range nums {
		num, _ := strconv.Atoi(strings.TrimSpace(numStr))
		result[i] = num
	}

	return result
}

// mengubah array []int{ 0,0,0 } jadi string `(0,0,0)`
func IntArrayToScores(input []int) string {
	// Convert each integer to a string
	strs := make([]string, len(input))
	for i, x := range input {
		strs[i] = strconv.Itoa(x)
	}

	// Calculate the maximum size of the resulting string
	size := 0
	for _, s := range strs {
		size += len(s) + 1 // Add 1 for the comma that will be added
	}
	size-- // Subtract 1 to account for the comma that comes after the last element

	// Create a buffer with the size of the resulting string
	buf := make([]byte, size+2) // Add 2 to account for the opening and closing parentheses

	// Add the opening parenthesis to the buffer
	buf[0] = '('

	// Iterate over the input array and append each element to the buffer
	last := len(strs) - 1
	offset := 1
	for i, s := range strs {
		// Convert the string to bytes and copy them to the buffer
		copy(buf[offset:], []byte(s))

		// Update the offset to point to the end of the current element
		offset += len(s)

		// Add a comma after the element if it's not the last one
		if i != last {
			buf[offset] = ','
			offset++
		}
	}

	// Add the closing parenthesis to the buffer
	buf[size+1] = ')'

	// Convert the buffer to a string and return it
	return string(buf)
}

// mengubah [][]int{{0,0,0},{0,0,0},{0,0,0}} jadi string ("(0,0,0)","(0,0,0)","(0,0,0)")
func IntArraysToScoresAndDuration(scores ...[]int) string {
	scoresStr := "("

	for i, score := range scores {
		scoresStr += "\"" + IntArrayToScores(score) + "\""

		if i != len(scores)-1 {
			scoresStr += ","
		}
	}

	scoresStr += ")"
	return scoresStr
}

func CheckmarksToBoolArray(input string) []bool {
	// Remove the parentheses from the input string
	input = strings.Trim(input, "()")

	// Split the remaining string into substrings
	substrings := strings.Split(input, ",")

	// Create a slice with the same length as the number of substrings
	output := make([]bool, len(substrings))

	// Iterate over the substrings and set the corresponding values in the output slice
	for i, substring := range substrings {
		// The substring is "f" for false or "t" for true
		switch substring {
		case "f":
			output[i] = false
		case "t":
			output[i] = true
		}
	}

	return output
}

func BoolArrayToCheckmarks(input []bool) string {
	// Create an empty string
	output := ""

	// Iterate over the input slice and append "f" for false and "t" for true
	for _, value := range input {
		if value {
			output += "t"
		} else {
			output += "f"
		}

		output += ","
	}

	// Remove the trailing comma and wrap the string in parentheses
	output = "(" + strings.TrimRight(output, ",") + ")"

	return output
}
