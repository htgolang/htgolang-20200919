package main

import (
	"fmt"
)

var slice1 = []int{108, 107, 105, 109, 103, 102}
var slice2 = []int{108, 107, 105, 109, 103, 102}

func main() {
	fmt.Println("var slice1 = []int{108, 107, 105, 109, 103, 102}")
	IdxMax, MaxValue := FindMax(slice1)
	fmt.Printf("max element is: %d\n", MaxValue)

	MaxToLast(IdxMax, MaxValue, len(slice1)-1, &slice1)
	fmt.Printf("move max element to last:\n%v\n", slice1)

	IdxMax, MaxValue = FindMax(slice1[:len(slice1)-1])
	fmt.Printf("second max element is: %d\n", MaxValue)

	MaxToLast(IdxMax, MaxValue, len(slice1)-2, &slice1)
	fmt.Printf("move second max element to second last:\n%v\n", slice1)

	//Bubble Sort
	for i := 0; i <= len(slice2)-1; i++ {
		IdxMax, MaxValue = FindMax(slice2[:len(slice2)-i])
		MaxToLast(IdxMax, MaxValue, len(slice2)-i-1, &slice2)
		fmt.Printf("Bubble Sort round %d: %v\n", i+1, slice2)
	}
	fmt.Println("Bubble Sort finish")
}
