package funcs

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		// if all the number shifted, sweep(nums, i) will return false
		if !sweep(nums, i) {
			return
		}
	}
}

func sweep(nums []int, passesDone int) bool {

	var firstIndex, secondIndex int = 0, 1
	doSwap := false

	// those already shifted is passesDone, minus them
	for secondIndex < len(nums)-passesDone {
		var firstNum = nums[firstIndex]
		var secondNum = nums[secondIndex]

		// shift big one backward
		if firstNum > secondNum {
			//nums[firstIndex] = secondNum
			//nums[secondIndex] = firstNum
			nums[firstIndex], nums[secondIndex] = secondNum, firstNum
			doSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return doSwap

}
