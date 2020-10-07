package main

import (
	"6-a_insert_sort/insertsort"
	"fmt"
)

var slice = []int{108, 107, 105, 109, 103, 102}

func main() {
	fmt.Println("var slice = []int{108, 107, 105, 109, 103, 102}")
	insertsort.InsertSort(&slice)
	fmt.Printf("After insert sort: %v\n", slice)
}
