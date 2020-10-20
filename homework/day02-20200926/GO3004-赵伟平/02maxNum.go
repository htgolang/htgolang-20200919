package main

import (
	"fmt"
)

func main() {
	var maxNum int = 0
	nums := [...]int{108, 107, 105, 109, 103, 102}

	for _, v := range nums {
		if v >= maxNum {
			maxNum = v
		}
	}

	fmt.Println(maxNum)
}
