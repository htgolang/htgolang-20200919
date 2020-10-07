package main

import (
	"6-a_insert_sort/insertsort"
	"fmt"
)

var slice = []int{108, 107, 105, 109, 103, 102}

const n int = 6

func main() {
	fmt.Println("var slice = []int{108, 107, 105, 109, 103, 102}")
	insertsort.InsertSort(&slice)
	fmt.Printf("sorted slice: %v\n", slice)
	RandMap := RandomMap(n)

	for idx, v := range slice {
		if _, ok := RandMap[v]; ok {
			RandMap[v] = idx
		}
	}

	fmt.Printf("total count of random number: %d\n", n)

	for k, v := range RandMap {
		fmt.Printf("random number: %d, index: %d\n", k, v)
	}
}
