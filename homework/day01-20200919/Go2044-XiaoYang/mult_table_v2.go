package main

import "fmt"

func main() {

	/*
		打印九九乘法表
	*/

	for {
		i := 0

		if i >= 9 {
			break
		} else {
			i++
			fmt.Println(i)
		}
	}

}
