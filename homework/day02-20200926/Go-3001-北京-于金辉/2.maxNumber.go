package main

import "fmt"

func main() {
	numbers := []int{108, 107, 105, 109, 103, 102}
	var maxNumber int

	for i := 0; i < len(numbers)-1; i++ {
		if maxNumber < numbers[i] {
			maxNumber = numbers[i]
		}
	}
	// 找到最大的数字
	fmt.Println(maxNumber)
}
