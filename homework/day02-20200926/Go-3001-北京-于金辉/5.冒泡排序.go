package main

import "fmt"

func main() {
	numbers := []int{108, 107, 105, 109, 103, 102}

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// [102 103 105 107 108 109]
	fmt.Println(numbers)
}
