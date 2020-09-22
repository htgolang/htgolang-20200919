package main

import "fmt"

func main() {

	/*
		打印九九乘法表
	*/

	first_loop := 0
	idx := 1

	for {
		// first loop
		first_loop += idx
		if first_loop <= 9 {
			for two_loop := 1; two_loop <= first_loop; two_loop++ {
				fmt.Printf("%d x %d = %d  ", first_loop, two_loop, first_loop*two_loop)

			}
			fmt.Println()
		} else {
			break
		}

	}

}
