package funcs

func GetMaxNum(numsArray []int) (int, int) {
	var maxNum, idxMax int
	for i := 0; i < len(numsArray); i++ {
		if numsArray[i] > maxNum {
			maxNum, idxMax = numsArray[i], i
		} else {
			continue
		}
	}
	return maxNum, idxMax
}
