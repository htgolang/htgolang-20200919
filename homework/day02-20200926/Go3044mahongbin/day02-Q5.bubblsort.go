package main

import (
	"fmt"
)

func main() {
	heights := []int{108, 107, 105, 109, 103, 102}
	for j := 0; j < len(heights)-1; j++ {
		fmt.Println("============================Round", j+1)
		for i := 0; i < len(heights)-1-j; i++ {
			fmt.Println("compare:", heights[i], heights[i+1])

			if heights[i] > heights[i+1] {
				fmt.Println("exchage:---> ", heights[i], heights[i+1])
				heights[i], heights[i+1] = heights[i+1], heights[i]
			}
			fmt.Println("result:", heights)
		}
		fmt.Println("Round result:", heights)
	}
}
