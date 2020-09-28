package funcs

func MoveToLast(numsArray []int) []int {
	var newNums = make([]int, len(numsArray))
	maxNum, idxMax := GetMaxNum(numsArray)
	for i := 0; i < len(numsArray); i++ {
		if i > idxMax {
			newNums[i-1] = numsArray[i]
		} else {
			newNums[i] = numsArray[i]
		}

	}
	//for i, v := range numsArray {
	//	if i != idxMax {
	//		newNums[i] = v
	//	} else {
	//		continue
	//	}
	//}
	newNums[len(newNums)-1] = maxNum
	return newNums
}
