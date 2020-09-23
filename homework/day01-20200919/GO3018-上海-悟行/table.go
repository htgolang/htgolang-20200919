package main

import "fmt"

func main() {
	fmt.Println("======打印九九乘法表======")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			//fmt.Printf("%v * %v = %2v\t", i, j, i*j)
			fmt.Printf("%d * %d = %2d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println("=======打印结束======")
}
