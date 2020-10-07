package main

import (
	"fmt"
)

func main() {
	var max int = 0
//	var num int = 0
	test := []int{108, 107, 105, 109, 103, 102}
	for i, _ := range test {
//		fmt.Printf("%d: %d\n",i, v)
		if max < test[i] {
			max = test[i]
		}
	}
	fmt.Printf("max num is %d\n", max)
}
