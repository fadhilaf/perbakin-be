package util

func CheckStage0SeriesScoresArray(arr []int) (bool, string) {
	if len(arr) != 11 {
		return false, "Panjang array harus 11"
	}

	sum := 0
	for _, val := range arr {
		sum += val
	}

	if sum > 10 {
		return false, "Total jumlah tembakan tidak boleh lebih dari 10"
	}

	return true, ""
}
