package main

import "fmt"

func main() {
	var nums = []int{108, 107, 105, 109, 103, 102}
	var MaxNum int
	for i := 0; i < len(nums); i++ {
		if nums[i] > MaxNum {
			MaxNum = nums[i]
		}
	}
	fmt.Println(MaxNum)

	
}
