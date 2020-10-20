package main

import (
	"fmt"
)

func main() {
	var test = []int{108, 107, 105, 109, 103, 102}
	for j := 0; j < len(test)-1; j++ {
		for i := 0; i < len(test)-1; i++ {
			if test[i] > test[i+1] {
				test[i], test[i+1] = test[i+1], test[i]
			}
		}
	}
	fmt.Printf("最大的数:%v\n", test[5])
	fmt.Printf("排序结果:%v\n", test)
}
