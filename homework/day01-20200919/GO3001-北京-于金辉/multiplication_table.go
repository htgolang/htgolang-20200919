package main

import "fmt"

func main() {
	fmt.Println("九九乘法表")
	for i := 1; i <= 9; i++ { // 列循环 1-9
		for j := 1; j <= i; j++ { // 行循环 1-n
			fmt.Printf("%dx%d=%-2d ", i, j, i*j)
		}
		fmt.Println()
	}
}
