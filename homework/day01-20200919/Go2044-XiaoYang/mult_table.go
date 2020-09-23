package main

import (
	"fmt"
)

func main() {

	/*
		打印九九乘法表
	*/

	for first_loop := 1; first_loop <= 9; first_loop++ {

		/*

			1,2,3,4,5,6,7,8

		*/
		for two_loop := 1; two_loop <= first_loop; two_loop++ {

			/*
				1
				12
				123
				1234
				1234...
			*/
			// fmt.Printf("第二层循环：%d\n", two_loop)
			fmt.Printf("%d*%d = %d ", two_loop, first_loop, first_loop*two_loop)

		}
		fmt.Println()
	}

}
