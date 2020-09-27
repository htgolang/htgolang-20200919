package main

import "fmt"

func main() {
	numbers := []int{108, 107, 105, 109, 103, 102}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]

		}
	}

	// 输出最大的一位
	fmt.Println(numbers[len(numbers)-1])
}
