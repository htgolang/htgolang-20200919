package main

import (
	"fmt"
)

func multiDemo(m int) {
	//打印乘法表
	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	multiDemo(9)
}
